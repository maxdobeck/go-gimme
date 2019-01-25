package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// parseClipboard finds stuff in the system clipboard
func parseClipboard() (map[string][]string, error) {
	found := make(map[string][]string)
	src := getClipboard()
	fmt.Println("Found these potential emails: ")
	words := strings.Fields(src)
	for _, word := range words {
		if strings.ContainsRune(word, 64) {
			found["email"] = append(found["email"], word)
		}
	}
	fmt.Println(found["email"])
	return found, nil
}

func parseFile(file string) (map[string][]string, error) {
	found := make(map[string][]string)
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	content, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	words := strings.Fields(string(content))
	for _, word := range words {
		if strings.ContainsRune(word, 64) {
			found["email"] = append(found["email"], word)
		}
	}
	fmt.Println(found["email"])
	return found, nil
}
