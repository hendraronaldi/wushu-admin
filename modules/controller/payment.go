package controller

import (
	"bytes"
	"encoding/base64"
	"os"
	"strings"
	"work/wushu-backend/chatbot/line"
	"work/wushu-backend/chatbot/line/messages"
	"work/wushu-backend/modules/connections"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func GetLineBotQR(c *gin.Context) {
	bucket := connections.FirebaseStorage()
	if data, err := connections.GetFileFirebaseStorage(bucket, "linebot-qr.png"); err != nil {
		c.JSON(400, gin.H{
			"response": "get linebot error",
		})
	} else {
		c.JSON(200, data)
	}

}

func PostPaymentConfirmation(c *gin.Context) {
	bucket := connections.FirebaseStorage()
	form, _ := c.MultipartForm()
	filename := strings.ToLower(form.Value["fullname"][0]) + "-" + form.Value["date"][0] + "." + form.Value["type"][0]
	files := form.Value["file"][0]
	idx := strings.Index(files, ";base64,")
	dec, err := base64.StdEncoding.DecodeString(files[idx+8:])
	if err != nil {
		panic(err)
	}

	res := bytes.NewReader(dec)
	if idx < 0 {
		c.JSON(400, gin.H{
			"response": "invalid image",
		})
	}

	if imgURL, err := connections.PostFileFirebaseStorage(bucket, filename, res); err != nil {
		c.JSON(400, gin.H{
			"response": "fail to save payment confirmation",
		})
	} else {
		var lineMessages []linebot.Message
		adminID := os.Getenv("line-admin-id")

		textMessage := messages.TextMessage(form.Value["fullname"][0] + "-" + form.Value["date"][0])
		imgMessage := messages.ImageMessage(imgURL)

		lineMessages = append(lineMessages, textMessage, imgMessage)

		if err := line.PushHandler(adminID, lineMessages); err != nil {
			c.JSON(400, gin.H{
				"response": "fail to send payment confirmation",
			})
		} else {
			c.JSON(200, gin.H{
				"response": "post payment confirmation success",
			})
		}
	}
}
