package messages

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func TextMessage(text string) *linebot.TextMessage {
	return linebot.NewTextMessage(text)
}
