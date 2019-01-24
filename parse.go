package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func parseClipboard(source *string) {
	fmt.Println("searching for emails in:\n", *source)
}

func parseFile(file string) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fp.Close()

	content, err := ioutil.ReadAll(fp)
	if err != nil {
		return err
	}
	fmt.Println("Searching for emails in: \n", file, content)
	return err
}
