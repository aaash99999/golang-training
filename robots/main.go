/*
Searches a URL for a given regular expression
*/
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
	arguments := os.Args[1:] // [1:] omits the program name, as we only want the args

	if len(arguments) != 2 {
		fmt.Printf("Usage: %s <URL> <regex to search for>\n", basename)
		os.Exit(1)
	}

	searchURL := arguments[0]
	searchString := arguments[1]

	res, err := http.Get(searchURL)

	if err != nil {
		log.Fatal(err)
	}

	pageBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Searching %s for '%s'\n", searchURL, searchString)
	matched, err := regexp.Match(searchString, pageBody)

	if err != nil {
		log.Fatal(err)
	}

	if matched {
		fmt.Printf("Success! Found: '%s'\n", searchString)
	} else {
		fmt.Printf("Fail! Did not find: '%s'\n", searchString)
	}

}
