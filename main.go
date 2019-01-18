package main

import "fmt"
import "flag"

func main() {
	emails := flag.Bool("emails", false, "Find all emails.  Defaults to false.")
	flag.Parse()
	if *emails {
		run(emails)
	} else {
		fmt.Println("gimme: Use -h to get help.")
	}
}

func run(emails *bool) {
	if *emails == true {
		findEmails()
	}
}
