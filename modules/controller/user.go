package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"work/wushu-backend/modules/connections"
	"work/wushu-backend/modules/model"

	"cloud.google.com/go/firestore"
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
	user.Status = 0
	conn := connections.FirebaseConnection()
	_, err := conn.Collection("users").Doc(user.Username).Set(context.Background(), user)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

func EditUser(c *gin.Context) {
	var user model.User
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid edit request",
		})
	} else {
		if _, err = FindUser(user.Username); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			conn := connections.FirebaseConnection()
			_, err := conn.Collection("users").Doc(user.Username).Set(context.Background(), user)
			if err != nil {
				c.JSON(400, gin.H{
					"response": "edit user error",
				})
			} else {
				c.JSON(200, gin.H{
					"response": "user is edited",
				})
			}
		}
	}
}

func ValidateUser(c *gin.Context) {
	var user model.User
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid validation request",
		})
	} else {
		if _, err = FindUser(user.Username); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			conn := connections.FirebaseConnection()
			_, err = conn.Collection("users").Doc(user.Username).Set(context.Background(), map[string]interface{}{
				"status": 1,
			}, firestore.MergeAll)
			if err != nil {
				c.JSON(400, gin.H{
					"response": "user validation error",
				})
			} else {
				c.JSON(200, gin.H{
					"response": "user is validated",
				})
			}
		}
	}
}

func RejectUser(c *gin.Context) {
	var user model.User
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid rejection request",
		})
	} else {
		if _, err = FindUser(user.Username); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			conn := connections.FirebaseConnection()
			_, err = conn.Collection("users").Doc(user.Username).Set(context.Background(), map[string]interface{}{
				"status": 2,
			}, firestore.MergeAll)
			if err != nil {
				c.JSON(400, gin.H{
					"response": "user rejection error",
				})
			} else {
				c.JSON(200, gin.H{
					"response": "user is rejected",
				})
			}
		}
	}
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
	var request model.Login
	var user map[string]interface{}
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid login request",
		})
	} else {
		if user, err = FindUser(request.Username); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			if request.Username == fmt.Sprint(user["Username"]) && request.Password == fmt.Sprint(user["Password"]) {
				c.JSON(200, user)
			} else {
				c.JSON(400, gin.H{
					"response": "wrong username or password",
				})
			}
		}
	}
}

func Logout(c *gin.Context) {

}
