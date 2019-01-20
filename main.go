package main

import (
	"flag"
)

func main() {
	emails := flag.Bool("emails", false, "Find all emails.  Defaults to false.")
	input := flag.String("input", "clipboard", "Specify the input file.  Defaults to your clipboard.")
	flag.Parse()

	if *input == "clipboard" {
		content := getClipboard()
		if *emails == true {
			parseClipboardForEmails(&content)
		}
	} else {
		validateFP(*input)
		if *emails == true {
			parseFileForEmails(*input)
		}
	}
}
