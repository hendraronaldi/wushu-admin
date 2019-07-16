package controller

import (
	"context"
	"encoding/json"
	"log"
	"work/wushu-backend/modules/connections"
	"work/wushu-backend/modules/model"

	"github.com/gin-gonic/gin"
)

func FindUser(username string) (map[string]interface{}, error) {
	conn := connections.FirebaseConnection()
	dsnap, err := conn.Collection("users").Doc(username).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := dsnap.Data()

	return user, nil
}

func AddUser(user model.User) error {
	conn := connections.FirebaseConnection()
	_, err := conn.Collection("users").Doc(user.Username).Set(context.Background(), user)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

func Register(c *gin.Context) {
	var newUser model.User
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&newUser); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid register request",
		})
	} else {
		if _, err = FindUser(newUser.Username); err == nil {
			c.JSON(400, gin.H{
				"response": "user existed already",
			})
		} else {
			// TODO: create new user
			if err = AddUser(newUser); err != nil {
				c.JSON(400, gin.H{
					"response": "add user error",
				})
			} else {
				c.JSON(200, gin.H{
					"response": "new user created",
				})
			}
		}
	}
}

func Login(c *gin.Context) {
	var request model.User
	// var user map[string]interface{}
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid login request",
		})
	} else {
		if _, err = FindUser(request.Username); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			// TODO: send user profile
		}
	}
}

func Logout(c *gin.Context) {

}
