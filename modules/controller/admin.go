package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"work/wushu-backend/modules/connections"
	"work/wushu-backend/modules/model"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func GetAllUser(c *gin.Context) {
	var users []map[string]interface{}
	var errs error

	conn := connections.FirebaseConnection()
	iter := conn.Collection("users").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			errs = err
			break
		}
		users = append(users, doc.Data())
	}

	if errs != nil {
		c.JSON(400, gin.H{
			"response": "get users error",
		})
	} else {
		c.JSON(200, gin.H{
			"response": "user is edited",
		})
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

func FindAdmin(username string) (map[string]interface{}, error) {
	conn := connections.FirebaseConnection()
	dsnap, err := conn.Collection("admin").Doc(username).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := dsnap.Data()

	return user, nil
}

func AdminLogin(c *gin.Context) {
	var request model.Login
	var user map[string]interface{}
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid login request",
		})
	} else {
		if user, err = FindAdmin(request.Username); err != nil {
			c.JSON(400, gin.H{
				"response": "admin not exist",
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
