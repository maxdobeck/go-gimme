package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
)

func getClipboard() string {
	content, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Problem reading from your clipboard\n: ", err)
	}
	return content
}

func validateFP(file string) {
	fmt.Println("attempting to open: ", file)
	fp, err := os.Open(file)
	if err != nil {
		fmt.Println("Problem opening file.  Verify file path or that file exists.  Erro:\n", err)
	}
	defer fp.Close()
}
