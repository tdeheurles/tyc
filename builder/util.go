package main

import (
	"fmt"
	"os"
)

func createDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}
}

func deleteDirectory(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}
