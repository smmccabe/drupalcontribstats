# drupalcontribstats

Simple go app for pulling and scanning through drupal modules and totalling up the unique number of contributors.

## Installation

### Linux

```
wget https://github.com/smmccabe/drupalcontribstats/releases/download/v0.1.0/drupalcontribstats-linux-amd64
sudo mv drupalcontribstats-linux-amd64 /usr/local/bin/drupalcontribstats
sudo chmod +x /usr/local/bin/drupalcontribstats
```

## Usage

```
drupalcontribstats commerce
```

You can pass in as many projects as you want
```
drupalcontribstats commerce commerce_pos
```

By default it takes the last year, but you can pass in custom starts and/or ends
```
drupalcontribstats --since 20190101 --until 20200101 commerce commerce_pos
```

Use a custom cache directory, you can always point this where you have many or all of the project cloned already
```
drupalcontribstats --cacheDir /home/me/dev drupal paragraphs
```

Pass in a list as a text file instead
```
drupalcontribstats --list mycustomlist.txt
```

Will process the custom list plus the ones provided by the command line
```
drupalcontribstats --list mycustomlist.txt paragraphs
```

Outputs the names of all the projects scanned and the contributors with their totals
```
drupalcontribstats commerce commerce_pos --verbose
```
