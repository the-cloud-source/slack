package main

import (
	"github.com/the-cloud-source/slack"
)

// Please change these values to suit your environment
const (
	token       = "your-api-token"
	channelName = "general"
)

func main() {
	api := slack.New(token)
	err := api.JoinChannel(channelName)
	if err != nil {
		panic(err)
	}
}
