package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Video is the data structure for the video content
type Video struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	Channel         string `json:"channel"`
	SummaryShortEn  string `json:"summary_short_en"`
	SummaryShortFr  string `json:"summary_short_fr"`
	SummaryMiddleEn string `json:"summary_middle_en"`
	SummaryMiddleFr string `json:"summary_middle_fr"`
	SummaryBigEn    string `json:"summary_big_en"`
	SummaryBigFr    string `json:"summary_big_fr"`
}

// VideoShort is the data structure which is defined in the content file
type VideoShort struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	SummaryShortEn  string `json:"summary_short_en"`
	SummaryShortFr  string `json:"summary_short_fr"`
	SummaryMiddleEn string `json:"summary_middle_en"`
	SummaryMiddleFr string `json:"summary_middle_fr"`
	SummaryBigEn    string `json:"summary_big_en"`
	SummaryBigFr    string `json:"summary_big_fr"`
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
