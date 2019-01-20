package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func parseClipboardForEmails(source *string) {
	fmt.Println("searching for emails in:\n", *source)
}

func parseFileForEmails(file string) {
	fp, openErr := os.Open(file)
	if openErr != nil {
		fmt.Println("Couldn't open file: \n", openErr)
	}
	content, err := ioutil.ReadAll(fp)
	if err != nil {
		fmt.Println("Something went wrong or we're at the end of the file. ", err)
	}
	fmt.Println("Searching for emails in: \n", string(content))
}
