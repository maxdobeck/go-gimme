package main

import (
	"os"
	"testing"
)

func TestParseInputFile(t *testing.T) {
	err := parseFile("nonexistent.txt")
	if err == nil {
		t.Fatal()
	}
}

func TestParseFile(t *testing.T) {
	_, createErr := os.Create("test.txt")
	if createErr != nil {
		t.Fatal()
	}
	openErr := parseFile("test.txt")
	if openErr != nil {
		t.Fatal(openErr)
	}
	os.Remove("test.txt")
}
