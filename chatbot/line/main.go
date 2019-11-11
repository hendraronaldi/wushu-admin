// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package line

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LineTP struct {
	bot *linebot.Client
}

var (
	baseURL = os.Getenv("APP_BASE_URL")
	botType = "line"
)

func App() *LineTP {
	var app *LineTP
	var err error
	//call API

	channelSecret := os.Getenv("line-channel-secret")
	channelAccessToken := os.Getenv("line-channel-access-token")
	app, err = CreateApp(
		channelSecret,
		channelAccessToken,
	)

	if err != nil {
		log.Fatal(err)
	}

	return app
}

func CreateApp(channelSecret, channelToken string) (*LineTP, error) {
	bot, err := linebot.New(
		channelSecret,
		channelToken,
	)

	if err != nil {
		return nil, err
	}
	return &LineTP{
		bot: bot,
	}, nil
}

func CallbackHandler(c *gin.Context) {
	app := App()

	events, err := app.bot.ParseRequest(c.Request)
	if err != nil {
		return
	}

	for _, event := range events {
		fmt.Println(event.Source.UserID)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				_, err := app.bot.ReplyMessage(event.ReplyToken, message).Do()
				if err != nil {
					log.Println("Quota err:", err)
				}

			case *linebot.ImageMessage:
				_, err := app.bot.ReplyMessage(event.ReplyToken, message).Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
			}
		}
	}
}
