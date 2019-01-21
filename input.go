package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func getClipboard() string {
	content, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Problem reading from your clipboard\n: ", err)
	}
	return content
}
