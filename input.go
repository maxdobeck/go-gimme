package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func parseClipboard() {
	content, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Problem reading from your clipboard\n: ", err)
	}
	fmt.Println("clipboard contents: \n", content)
}

func parseFile(path string) {
	fmt.Println(path)
}
