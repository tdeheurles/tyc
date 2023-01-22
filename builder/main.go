package main

// Expect this code to generate the hugo content from data

func main() {
	contentData := getContentData()

	deleteDirectory(contentPath)
	createDirectory(channelPath)
	createDirectory(videoPath)

	for _, channel := range contentData.Channels {
		createChannelContent(channel)

		for _, video := range channel.Videos {
			createVideoFile(video, channel)
		}
	}
}
