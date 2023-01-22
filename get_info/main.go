package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/youtube/v3"
)

func main() {
	client, err := youtube.NewService(context.Background())
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	channelName := "Kurzgesagt"
	channel, err := getChannelByName(client, channelName)
	if err != nil {
		log.Fatalf("Error getting channel: %v", err)
	}

	fmt.Println("Channel ID:", channel.Id)
}

func getChannelByName(service *youtube.Service, name string) (*youtube.Channel, error) {
	call := service.Search.List([]string{"id"}).Q(name).Type("channel")

	response, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("error calling youTube API: %v", err)
	}

	if len(response.Items) == 0 {
		return nil, fmt.Errorf("no channel found with name %q", name)
	}

	channelId := response.Items[0].Id.ChannelId
	channel, err := service.Channels.List([]string{"snippet"}).Id(channelId).Do()
	if err != nil {
		return nil, fmt.Errorf("error getting channel details: %v", err)
	}
	return channel.Items[0], nil
}
