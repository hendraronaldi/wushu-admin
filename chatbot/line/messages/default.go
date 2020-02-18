package messages

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

var defaultMenu = []string{"Register"}
var homeMenu = []string{"Proof of Payment"}

func BeginMessage() []linebot.Message {
	var template []linebot.Message
	buttonMessage := ButtonMessage(defaultMenu)

	template = append(template, buttonMessage)
	return template
}

func HomeMessage() []linebot.Message {
	var template []linebot.Message
	buttonMessage := ButtonMessage(homeMenu)

	template = append(template, buttonMessage)
	return template
}
