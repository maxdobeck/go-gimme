package main

import (
	"flag"
	"fmt"
)

func main() {
	emails := flag.Bool("emails", false, "Find all emails.  Defaults to false.")
	input := flag.String("f", "clipboard", "Specify the input file.  Defaults to your clipboard.")
	flag.Parse()

	// decide to get data from clipboard or a supplied file
	if *input != "clipboard" {
		if *emails == true {
			err := parseFile(*input)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		if *emails == true {
			found, err := parseClipboard()
			check(err)
			fmt.Println(found["email"])
		}
	}

	help(emails)
}

func help(emails *bool) {
	if *emails == false && len(flag.Args()) < 1 {
		fmt.Println("gimme: gets you stuff from data.  use gimme -h for help.")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
