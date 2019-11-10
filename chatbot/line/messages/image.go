package messages

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func ImageMessage(url string) *linebot.ImageMessage {
	return linebot.NewImageMessage(url, url)
}
