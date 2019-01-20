package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func parseInput() {
	content, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Problem reading from your clipboard\n: ", err)
	}
	fmt.Println("clipboard contents: \n", content)
}
