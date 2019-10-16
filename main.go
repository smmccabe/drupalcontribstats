package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"github.com/btubbs/datetime"
)

func main() {
	var contributors = make(map[string]int)

	sinceFlag := flag.String("since", "now - 1 year", "a to start from in any ISO 8601 format")
	untilFlag := flag.String("until", "now", "a to count until in any ISO 8601 format")
	cacheDir := flag.String("cacheDir", "/tmp/drupalcontribstats", "folder to store repositories, can already be populated")
	listFile := flag.String("list", "", "a file with a newline seperated list of drupal project to scan")
	verbose := flag.Bool("verbose", false, "Display extra debug information")

	flag.Parse()

	var err error
	var since time.Time
	if *sinceFlag == "now - 1 year" {
		since = time.Now().AddDate(-1, 0, 0)
	} else {
		since, err = datetime.Parse(*sinceFlag, time.UTC)
		if err != nil {
			log.Fatal(err)
		}
	}

	var until time.Time
	if *untilFlag == "now" {
		until = time.Now()
	} else {
		until, err = datetime.Parse(*untilFlag, time.UTC)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *listFile != "" {
		file, err := os.Open(*listFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			tallyRepo(text, contributors, since, until, *cacheDir, *verbose)
		}
	}

	args := flag.Args()
	if len(args) > 0 {
		for _, arg := range args {
			tallyRepo(arg, contributors, since, until, *cacheDir, *verbose)
		}
	}

	if *verbose {
		for contributor, contributions := range contributors {
			fmt.Println(contributor+":", contributions)
		}
	}
	fmt.Println(len(contributors))
}

func tallyRepo(project string, contributors map[string]int, since time.Time, until time.Time, cacheDir string, verbose bool) {
	re := regexp.MustCompile(`Issue #(.*) by ([^:]*)`)
	url := "https://git.drupalcode.org/project/" + project + ".git"
	path := path.Join(cacheDir, project)

	var r *git.Repository
	var err error
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if verbose {
			fmt.Println("Opening existing " + project + "...")
		}
		r, err = git.PlainOpen(path)
	} else {
		if verbose {
			fmt.Println("Cloning " + project + "...")
		}
		r, err = git.PlainClone(path, false, &git.CloneOptions{
			URL: url,
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	ref, err := r.Head()

	if err != nil {
		log.Fatal(err)
	}

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})

	if err != nil {
		log.Fatal(err)
	}

	err = cIter.ForEach(func(c *object.Commit) error {
		var substring = re.FindStringSubmatch(c.Message)
		if len(substring) < 3 {
			return nil
		}
		var split = strings.Split(substring[2], ", ")

		for _, each := range split {
			contributors[each] = contributors[each] + 1
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
