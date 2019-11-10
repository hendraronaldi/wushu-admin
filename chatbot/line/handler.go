package line

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func ReplyHandler(text string) {

}

func PushHandler(id string, messages []linebot.Message) error {
	app := App()
	_, err := app.bot.PushMessage(
		id,
		messages...,
	).Do()

	return err
}
