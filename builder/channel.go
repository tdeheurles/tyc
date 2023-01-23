package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Channel struct {
	Title                   string  `json:"title"`
	summary_one_sentence_en string  `json:"summary_one_sentence_en"`
	Videos                  []Video `json:"videos"`
}

func createChannelContent(channel Channel) {
	fileName := strings.ReplaceAll(channel.Title, " ", "-")
	f, err := os.Create(channelPath + fileName + ".md")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	content, err := json.MarshalIndent(channel, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f.Write([]byte("---\n"))
	f.Write([]byte(content))
	f.Write([]byte("\n"))
	f.Write([]byte("---\n"))
	f.Write([]byte("\n"))
	f.Sync()
}
