package main

import (
	"github.com/atotto/clipboard"
	"testing"
)

func TestClipboardIsReachable(t *testing.T) {
	err := clipboard.WriteAll("This should be in the OS clipboard")
	if err != nil {
		t.Fatal(err)
	}
	content := parseClipboard()
	if content != "This should be in the OS clipboard" {
		t.Fatalf("String in clipboard doesn't match the expected result")
	}
}
