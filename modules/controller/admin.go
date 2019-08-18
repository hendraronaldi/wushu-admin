package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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
		c.JSON(200, users)
	}
}

func GetUserByStatus(c *gin.Context) {
	var users []map[string]interface{}
	var errs error
	var codeStatus int
	status := c.Param("status")
	if strings.ToLower(status) == "verified" {
		codeStatus = 1
	} else if strings.ToLower(status) == "rejected" {
		codeStatus = 2
	} else {
		codeStatus = 0
	}

	conn := connections.FirebaseConnection()
	iter := conn.Collection("users").Where("Status", "==", codeStatus).Documents(context.Background())
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
		c.JSON(200, users)
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
		user.Email = strings.ToLower(user.Email)
		if _, err = FindUser(user.Email); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			conn := connections.FirebaseConnection()
			_, err = conn.Collection("users").Doc(user.Email).Set(context.Background(), map[string]interface{}{
				"Status": 1,
			}, firestore.MergeAll)
			if err != nil {
				c.JSON(400, gin.H{
					"response": "user validation error",
				})
			} else {
				user.Status = 1
				c.JSON(200, user)
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
		user.Email = strings.ToLower(user.Email)
		if _, err = FindUser(user.Email); err != nil {
			c.JSON(400, gin.H{
				"response": "user not exist",
			})
		} else {
			conn := connections.FirebaseConnection()
			_, err = conn.Collection("users").Doc(user.Email).Set(context.Background(), map[string]interface{}{
				"Status": 2,
			}, firestore.MergeAll)
			if err != nil {
				c.JSON(400, gin.H{
					"response": "user rejection error",
				})
			} else {
				user.Status = 2
				c.JSON(200, user)
			}
		}
	}
}

func FindAdmin(email string) (map[string]interface{}, error) {
	conn := connections.FirebaseConnection()
	dsnap, err := conn.Collection("admin").Doc(email).Get(context.Background())
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
		request.Email = strings.ToLower(request.Email)
		if user, err = FindAdmin(request.Email); err != nil {
			c.JSON(400, gin.H{
				"response": "admin not exist",
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
