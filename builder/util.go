package main

import (
	"fmt"
	"os"
)

// createDirectory creates a directory if it does not exist
func createDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}
}

// deleteDirectory deletes a directory
func deleteDirectory(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}
