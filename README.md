# drupalcontribstats

Simple go app for pulling and scanning through drupal modules and totalling up the unique number of contributors. Right now it is fairly hardcoded and doesn't take many parameters

## Usage

```
drupalcontribstats commerce

# You can pass in as many projects as you want
drupalcontribstats commerce commerce_pos

# by default it takes the last year, but you can pass in custom starts and/or ends
drupalcontribstats --since 20190101 --until 20200101 commerce commerce_pos

# a custom cache directory, you can always point this where you have many or all of the project cloned already
drupalcontribstats --cacheDir /home/me/dev drupal paragraphs

# pass in a list as a text file instead
drupalcontribstats --list mycustomlist.txt

# Will process the custom list plus the ones provided by the command line
drupalcontribstats --list mycustomlist.txt paragraphs

# Outputs the names of all the projects scanned and the contributors with their totals
drupalcontribstats commerce commerce_pos --verbose
```
