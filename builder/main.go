package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Expect this code to generate the hugo content from data

type Video struct {
	Id      string `yaml:"id" json:"id"`
	Summary string `yaml:"summary" json:"summary"`
	Title   string `yaml:"title" json:"title"`
}

type Channel struct {
	Title   string  `yaml:"title" json:"title"`
	Summary string  `yaml:"summary" json:"summary"`
	Videos  []Video `yaml:"videos" json:"videos"`
}

type Data struct {
	Channels []Channel `yaml:"channels" json:"channels"`
}

func main() {
	data := getFromJson()

	fmt.Printf("%d channels to build\n\n", len(data.Channels))

	videoDirectory := "../hugo/content/videos/"
	if _, err := os.Stat(videoDirectory); os.IsNotExist(err) {
		os.MkdirAll(videoDirectory, 0700) // Create your file
	}

	for _, channel := range data.Channels {
		for _, video := range channel.Videos {

			f, err := os.Create(videoDirectory + video.Id + ".md")
			if err != nil {
				panic(err)
			}
			defer f.Close()

			f.Write([]byte("---\n"))
			f.Write([]byte(fmt.Sprintf("id: %s\n", video.Id)))
			f.Write([]byte(fmt.Sprintf("summary: \"%s\"\n", video.Summary)))
			f.Write([]byte(fmt.Sprintf("title: \"%s\"\n", video.Title)))
			f.Write([]byte("---\n"))
			f.Write([]byte(fmt.Sprintf("{{< youtube %s >}}\n", video.Id)))
			f.Sync()

			fmt.Printf("video %s - %s\n", video.Title, video.Summary)
		}
	}
}

func getFromJson() Data {
	// Read file
	content, err := os.ReadFile("../hugo/data/all.json")
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
