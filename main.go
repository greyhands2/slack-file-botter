package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xxxxxxxx")
	os.Setenv("CHANNEL_ID", "xxxxxxxx")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"grokking_algorithms_2016.pdf"}

	//using a for loop
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("%s[1] \n", err)
			return
		}

		fmt.Printf("Name: %s[1] , URL: %s[2]", file.Name, file.URL)
	}
}
