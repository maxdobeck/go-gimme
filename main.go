package main

import (
	"flag"
	"fmt"
)

func main() {
	emails := flag.Bool("emails", false, "Find all emails.  Defaults to false.")
	input := flag.String("f", "clipboard", "Specify the input file.  Defaults to your clipboard.")
	flag.Parse()

	if *input == "clipboard" {
		content := getClipboard()
		if *emails == true {
			parseClipboard(&content)
		}
	} else {
		if *emails == true {
			err := parseFile(*input)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
