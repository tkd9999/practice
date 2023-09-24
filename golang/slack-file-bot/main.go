package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5934599514294-5938448706917-NWLspNJyy7fiKaV9iX7xnGYB")
	os.Setenv("CHANNEL_ID", "C05U1RWKTH7")
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
