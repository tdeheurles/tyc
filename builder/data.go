package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Data is the data structure for the JSON file
type Data struct {
	Channels []Channel `json:"channels"`
}

// Read the content data from the JSON file
func getContentData() Data {
	// Read file
	content, err := os.ReadFile(contentDataPath)
	if err != nil {
		panic(fmt.Errorf("error reading file: %s", err))
	}

	// Unmarshal JSON data
	data := Data{}
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		panic(err)
	}
	return data
}
