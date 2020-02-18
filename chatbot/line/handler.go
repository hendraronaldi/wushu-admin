package line

import (
	"fmt"
	"os"
	"strings"
	"time"
	"work/wushu-backend/chatbot/line/controller"
	"work/wushu-backend/chatbot/line/messages"
	"work/wushu-backend/chatbot/line/model"

	"github.com/line/line-bot-sdk-go/linebot"
)

var adminID = os.Getenv("line-admin-id")

func ReplyHandler(id string, m linebot.Message) []linebot.Message {
	var riveReply string
	var botPushMessage []linebot.Message
	var botReply []linebot.Message
	var menu []linebot.Message

	// User bot
	if id != adminID {
		user, err := controller.FindLineUser(id)
		if err != nil || user == nil || fmt.Sprint(user["ID"]) != id {
			LoadContext(id, "new")
			menu = messages.BeginMessage()
			switch message := m.(type) {
			case *linebot.TextMessage:
				riveReply = GetBotReply(message.Text, id, "new")
				fmt.Println("line 32", riveReply)
				if message.Text == riveReply {
					var confirmation []map[string]string
					yes := make(map[string]string)
					no := make(map[string]string)

					yes["Yes"] = "yes\nregistration\n" + id + "\n" + riveReply
					no["No"] = "no\nregistration\n" + id + "\n" + riveReply
					confirmation = append(confirmation, yes, no)

					botReply = append(botReply, messages.TextMessage("Your name is: "+riveReply))
					botReply = append(botReply, messages.ConfirmCustomMessage(confirmation))
				} else {
					botReply = append(botReply, messages.TextMessage(riveReply))
					if riveReply == "hello, welcome to Teratai Putih" {
						botReply = append(botReply, menu...)
					}
				}
			default:
				riveReply = GetBotReply("home", id, "new")
				fmt.Println("line 52", riveReply)
				botReply = append(botReply, messages.TextMessage(riveReply))
				botReply = append(botReply, menu...)
			}
			return botReply

		} else if fmt.Sprint(user["Status"]) == "false" {
			LoadContext(id, "new")
			botReply = append(botReply, messages.TextMessage("Please wait for the user verification"))
			return botReply
		}

		LoadContext(id, "registered")
		menu = messages.HomeMessage()
		switch message := m.(type) {
		case *linebot.TextMessage:
			if strings.ToLower(message.Text) == "register" {
				botReply = append(botReply, messages.TextMessage("Account already registered"))
				botReply = append(botReply, menu...)
			} else {
				riveReply = GetBotReply(message.Text, id, "registered")
				fmt.Println("line 73", riveReply)
				botReply = append(botReply, messages.TextMessage(riveReply))
				if riveReply == "hello, welcome to Teratai Putih" {
					botReply = append(botReply, menu...)
				}
			}

		case *linebot.ImageMessage:
			riveReply = GetBotReply("yes proof of payment", id, "registered")
			fmt.Println("line 81", riveReply)
			if riveReply == "proof of payment" {
				var confirmation []map[string]string
				yes := make(map[string]string)
				no := make(map[string]string)

				yes["Yes"] = "yes\nproof of payment\n" + id + "\n" + message.OriginalContentURL
				no["No"] = "no\nproof of payment\n" + id + "\n" + message.OriginalContentURL
				confirmation = append(confirmation, yes, no)

				textMessage := messages.TextMessage(fmt.Sprint(user["Name"]) + "'s proof of payment")
				botPushMessage = append(botPushMessage, textMessage, message, messages.ConfirmCustomMessage(confirmation))
				err := PushHandler(adminID, botPushMessage)
				if err != nil {
					GetBotReply("payment", id, "registered")
					botReply = append(botReply, messages.TextMessage("Fail to send proof of payment, please send it again"))
				} else {
					botReply = append(botReply, messages.TextMessage("Your proof of payment has been sent, please wait for the confirmation"))
				}
			} else {
				botReply = append(botReply, messages.TextMessage(riveReply))
				botReply = append(botReply, menu...)
			}
		}
		return botReply
	}

	// Admin bot
	switch message := m.(type) {
	case *linebot.TextMessage:
		if strings.HasPrefix(message.Text, "yes\n") || strings.HasPrefix(message.Text, "no\n") {
			confirmationDetails := strings.Split(message.Text, "\n")
			user, err := controller.FindLineUser(confirmationDetails[2])
			if strings.HasPrefix(message.Text, "yes\n") {
				if confirmationDetails[1] == "registration" {
					if err != nil {
						botPushMessage = append(botPushMessage, messages.TextMessage("Verification account failed"))
					} else {
						// update user status after confirmation
						var confirmUser model.LineUser
						confirmUser.ID = fmt.Sprint(user["ID"])
						confirmUser.Name = fmt.Sprint(user["Name"])
						confirmUser.Status = true
						err = controller.EditLineUser(confirmUser)
						if err != nil {
							botPushMessage = append(botPushMessage, messages.TextMessage("Verification account failed"))
						} else {
							botPushMessage = append(botPushMessage, messages.TextMessage("Your account has been verified"))
						}
					}
				} else {
					// save image to firebase storage
					t := time.Now().Format("2006-01-02 15:04:05")
					if isSavedProofOfPayment := controller.SaveProofOfPayment(confirmationDetails[3], t+" "+fmt.Sprint(user["Name"])+" "+fmt.Sprint(user["ID"])); isSavedProofOfPayment != 0 {
						botPushMessage = append(botPushMessage, messages.TextMessage("Your proof of payment has been rejected"))
					} else {
						botPushMessage = append(botPushMessage, messages.TextMessage("Your proof of payment has been verified"))
					}
				}
			} else {
				if confirmationDetails[1] == "registration" {
					botPushMessage = append(botPushMessage, messages.TextMessage("Verification account failed"))
				} else {
					botPushMessage = append(botPushMessage, messages.TextMessage("Your proof of payment has been rejected"))
				}
			}
			err = PushHandler(confirmationDetails[2], botPushMessage)
			if err != nil {
				botReply = append(botReply, messages.TextMessage("Failed to send verification to ID: "+confirmationDetails[2]+"\nProcess: "+confirmationDetails[1]))
			} else {
				botReply = append(botReply, messages.TextMessage("Verification success to ID: "+confirmationDetails[2]+"\nProcess: "+confirmationDetails[1]))
			}
		} else {
			botReply = append(botReply, messages.TextMessage("Unknown Action"))
		}
	default:
		botReply = append(botReply, messages.TextMessage("Unknown Action"))
	}
	return botReply
}

func PushHandler(id string, messages []linebot.Message) error {
	app := App()
	_, err := app.bot.PushMessage(
		id,
		messages...,
	).Do()

	if err != nil {
		return err
	}
	return nil
}
