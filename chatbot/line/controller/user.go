package controller

import (
	"context"
	"log"
	"strings"
	"work/wushu-backend/chatbot/line/model"
	"work/wushu-backend/modules/connections"
)

// Line
func FindLineUser(id string) (map[string]interface{}, error) {
	conn := connections.FirebaseConnection()
	dsnap, err := conn.Collection("line-user").Doc(id).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := dsnap.Data()

	return user, nil
}

func AddLineUser(user model.LineUser) error {
	user.Name = strings.ToLower(user.Name)
	user.Status = false
	conn := connections.FirebaseConnection()
	_, err := conn.Collection("line-user").Doc(user.ID).Set(context.Background(), user)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

func EditLineUser(user model.LineUser) error {
	var err error
	if _, err = FindLineUser(user.ID); err != nil {
		return err
	} else {
		conn := connections.FirebaseConnection()
		_, err := conn.Collection("line-user").Doc(user.ID).Set(context.Background(), user)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteLineUser(id string) error {
	var err error
	if _, err = FindLineUser(id); err != nil {
		return err
	} else {
		conn := connections.FirebaseConnection()
		_, err := conn.Collection("line-user").Doc(id).Delete(context.Background())
		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
			return err
		}
		return nil
	}
}
