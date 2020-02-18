package messages

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func ButtonMessage(text string, options []string) *linebot.TemplateMessage {
	var actions []linebot.TemplateAction
	for _, option := range options {
		actions = append(actions, linebot.NewMessageTemplateAction(option, option))
	}

	template := linebot.NewButtonsTemplate(
		"",
		"",
		text,
		actions...,
	)

	return linebot.NewTemplateMessage("button", template)
}

func ConfirmCustomMessage(text string, options []map[string]string) *linebot.TemplateMessage {
	var actions []linebot.TemplateAction
	for _, option := range options {
		for key, val := range option {
			actions = append(actions, linebot.NewMessageTemplateAction(key, val))
		}
	}

	template := linebot.NewButtonsTemplate(
		"",
		"",
		text,
		actions...,
	)

	return linebot.NewTemplateMessage("button", template)
}
