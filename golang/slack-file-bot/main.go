package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "bot-token")
	os.Setenv("CHANNEL_ID", "channel-id")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"index.html"}

	for i := 0; i < len(fileArr); i++ {
		parmas := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(parmas)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Nama: %s, URL:%s\n", file.Name, file.URL)
	}
}
