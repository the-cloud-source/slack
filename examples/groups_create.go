package main

import (
	"github.com/the-cloud-source/slack"
)

// Please change these values to suit your environment
const (
	token     = "your-api-token"
	groupName = "create-group"
)

func main() {
	api := slack.New(token)
	err := api.CreateGroup(groupName)
	if err != nil {
		panic(err)
	}
}
