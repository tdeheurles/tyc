package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Channels []Channel `json:"channels"`
}

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
