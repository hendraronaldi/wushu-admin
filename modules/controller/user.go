package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"work/wushu-backend/modules/connections"
	"work/wushu-backend/modules/model"

	"github.com/gin-gonic/gin"
)

func FindUser(email string) (map[string]interface{}, error) {
	conn := connections.FirebaseConnection()
	dsnap, err := conn.Collection("users").Doc(email).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := dsnap.Data()

	return user, nil
}

func AddUser(user model.User) error {
	user.Status = 0
	conn := connections.FirebaseConnection()
	_, err := conn.Collection("users").Doc(user.Email).Set(context.Background(), user)
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
		user.Email = strings.ToLower(user.Email)
		if _, err = FindUser(user.Email); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			conn := connections.FirebaseConnection()
			_, err := conn.Collection("users").Doc(user.Email).Set(context.Background(), user)
			if err != nil {
				c.JSON(400, gin.H{
					"response": "edit user error",
				})
			} else {
				c.JSON(200, user)
			}
		}
	}
}

func DeleteUser(c *gin.Context) {
	var user model.User
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid deletion request",
		})
	} else {
		user.Email = strings.ToLower(user.Email)
		if _, err = FindUser(user.Email); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			conn := connections.FirebaseConnection()
			_, err := conn.Collection("users").Doc(user.Email).Delete(context.Background())
			if err != nil {
				// Handle any errors in an appropriate way, such as returning them.
				log.Printf("An error has occurred: %s", err)
			}
			if err != nil {
				c.JSON(400, gin.H{
					"response": "user deletion error",
				})
			} else {
				c.JSON(200, gin.H{
					"response": "user is deleted",
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
		newUser.Email = strings.ToLower(newUser.Email)
		if _, err = FindUser(newUser.Email); err == nil {
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
		request.Email = strings.ToLower(request.Email)
		if user, err = FindUser(request.Email); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			if request.Email == fmt.Sprint(user["Email"]) && request.Password == fmt.Sprint(user["Password"]) {
				c.JSON(200, user)
			} else {
				c.JSON(400, gin.H{
					"response": "wrong email or password",
				})
			}
		}
	}
}

func Logout(c *gin.Context) {

}
