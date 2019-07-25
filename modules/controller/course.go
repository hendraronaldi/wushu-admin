package controller

import (
	"context"
	"encoding/json"
	"work/wushu-backend/modules/connections"
	"work/wushu-backend/modules/model"

	"github.com/gin-gonic/gin"
)

func GetCourse(c *gin.Context) {
	courseID := c.Param("id")
	var err error

	conn := connections.FirebaseTutorialConnection()
	dsnap, err := conn.Collection("courses").Doc(courseID).Get(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"response": "get course error",
		})
	} else {
		course := dsnap.Data()
		c.JSON(200, course)
	}
}

func PostCourse(c *gin.Context) {
	var course model.Course
	var err error

	if json.NewDecoder(c.Request.Body).Decode(&course); err != nil {
		c.JSON(400, gin.H{
			"response": "invalid course request",
		})
	} else {
		conn := connections.FirebaseTutorialConnection()
		_, err := conn.Collection("courses").Doc(course.ID).Set(context.Background(), course)
		if err != nil {
			c.JSON(400, gin.H{
				"response": "add courses error",
			})
		} else {
			c.JSON(200, course)
		}
	}
}
