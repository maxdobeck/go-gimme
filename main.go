package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
)

func main() {
	emails := flag.Bool("emails", false, "Find all emails.  Defaults to false.")
	input := flag.String("f", "clipboard", "Specify the input file.  Defaults to your clipboard.")
	flag.Parse()

	// decide to get data from clipboard or a supplied file
	if *input != "clipboard" {
		if *emails == true {
			found, err := parseFile(*input)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(found["email"])
		}
	} else {
		if *emails == true {
			found, err := parseClipboard()
			check(err)
			cpEmails(found["email"])
		}
	}

	help(emails)
}

func help(emails *bool) {
	if *emails == false && len(flag.Args()) < 1 {
		fmt.Println("gimme: gets you stuff from data.  use gimme -h for help.")
	}
}

func cpEmails(data []string) {
	var cp string
	for _, email := range data {
		cp = cp + email + "\n"
	}
	err := clipboard.WriteAll(cp)
	if err != nil {
		panic(err)
	}
	fmt.Println(cp)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
