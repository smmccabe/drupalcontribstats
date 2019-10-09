package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"bufio"
	"log"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	var contributors = make(map[string]int)

	tallyRepo("https://git.drupalcode.org/project/commerce.git", contributors)

	file, err := os.Open("list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		tallyRepo("https://git.drupalcode.org/project/commerce_" + text + ".git", contributors)
	}
	
	spew.Dump(contributors)
	fmt.Println(len(contributors))
}

func tallyRepo(url string, contributors map[string]int) {
	re := regexp.MustCompile(`Issue #(.*) by ([^:]*)`)
	folder := url[strings.LastIndex(url, "/")+1:]
	path := "/tmp/git/" + folder

	var r *git.Repository
	var err error
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		r, err = git.PlainOpen(path)
	} else {
		r, err = git.PlainClone(path, false, &git.CloneOptions{
			URL: url,
		})
	}

	if err != nil {
		log.Fatal(err)
	}

	ref, err := r.Head()

	since := time.Date(2018, time.October, 8, 0, 0, 0, 0, time.UTC)
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since})

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
