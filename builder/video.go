package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Video is the data structure for the video content
type Video struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Channel       string `json:"channel"`
	SummaryShort  string `json:"summary_short"`
	SummaryMiddle string `json:"summary_middle"`
	SummaryBig    string `json:"summary_big"`
}

// VideoShort is the data structure which is defined in the content file
type VideoShort struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	SummaryShort  string `json:"summary_short"`
	SummaryMiddle string `json:"summary_middle"`
	SummaryBig    string `json:"summary_big"`
}

// createVideoFile creates a video content file
func createVideoFile(video Video, channel Channel) {

	channelTitle := strings.ReplaceAll(channel.Title, " ", "_")
	channelTitle = strings.ToLower(channelTitle)
	video.Channel = channelTitle

	f, err := os.Create(videoPath + video.Id + ".md")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	content, err := json.MarshalIndent(video, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	f.Write([]byte("---\n"))
	f.Write([]byte(content))
	f.Write([]byte("\n"))
	f.Write([]byte("---\n"))
	f.Write([]byte("\n"))
	f.Write([]byte(fmt.Sprintf("{{< youtube %s >}}\n", video.Id)))
	f.Sync()
}
