package messages

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// var defaultMenu = []string{"Register"}
// var homeMenu = []string{"Proof of Payment"}

func BeginMessage() []linebot.Message {
	defaultMenu := []string{"Register"}
	var template []linebot.Message
	buttonMessage := ButtonMessage("Home", defaultMenu)

	template = append(template, buttonMessage)
	return template
}

func HomeMessage() []linebot.Message {
	homeMenu := []string{"Proof of Payment"}
	var template []linebot.Message
	buttonMessage := ButtonMessage("Home", homeMenu)

	template = append(template, buttonMessage)
	return template
}
