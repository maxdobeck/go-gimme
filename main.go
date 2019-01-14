package main

import "fmt"
import "flag"

func main() {
	fmt.Println("welcome to gimme.  Use -h to get help.")
	emails := flag.Bool("emails", false, "Find all emails.  Defaults to false.")
	flag.Parse()
	run(emails)
}

func run(emails *bool) {
	if *emails == true {
		findEmails()
	}
}
