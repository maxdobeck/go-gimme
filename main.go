package main

import "fmt"
import "flag"

func main() {
	emails := flag.Bool("emails", false, "Find all emails.  Defaults to false.")
	input := flag.String("input", "clipboard", "Specify the input file.  Defaults to your clipboard.")
	flag.Parse()

	if *input == "clipboard" {
		parseClipboard()
	} else if *input != "clipboard" {
		parseFile(*input)
	} else if *emails == true {
		findEmails()
	} else {
		fmt.Println("gimme: Use -h to get help.")
	}
}
