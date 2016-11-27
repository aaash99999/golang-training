package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	basename := filepath.Base(os.Args[0])
	site := os.Args[1:] // don't include program name
	if len(site) != 2 {
		fmt.Printf("Usage: %s <URL> <regex to search for>\n", basename)
		os.Exit(1)
	}
	// site[0] is the URL to check
	// site[1] is the regex to search for

	searchstring := site[1]

	res, err := http.Get(site[0])

	if err != nil {
		log.Fatal(err)
	}

	pagebody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Searching %s for '%s'\n", site[0], searchstring)
	matched, err := regexp.Match(searchstring, pagebody)

	if err != nil {
		log.Fatal(err)
	}

	if matched {
		fmt.Printf("Success! Found: '%s'\n", searchstring)
	} else {
		fmt.Printf("Fail! Did not find: '%s'\n", searchstring)
	}

}
