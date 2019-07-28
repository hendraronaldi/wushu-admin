package controller

import (
	"context"
	"encoding/json"
	"work/wushu-backend/modules/connections"
	"work/wushu-backend/modules/model"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func GetCourseClass(c *gin.Context) {
	var errs error
	var classes []map[string]interface{}

	conn := connections.FirebaseTutorialConnection()
	iter := conn.Collection("courses").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			errs = err
			break
		}
		classes = append(classes, doc.Data())
	}

	if errs != nil {
		c.JSON(400, gin.H{
			"response": "get course class error",
		})
	} else {
		c.JSON(200, classes)
	}
}

func GetCourseCategory(c *gin.Context) {
	courseClass := c.Param("class")
	courseCategory := c.Param("category")

	var errs error
	var courses []string

	conn := connections.FirebaseTutorialConnection()
	iter := conn.Collection("courses/" + courseClass + "/" + courseCategory).DocumentRefs(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			errs = err
			break
		}
		courses = append(courses, doc.ID)
	}

	if errs != nil {
		c.JSON(400, gin.H{
			"response": "get course class error",
		})
	} else {
		c.JSON(200, courses)
	}
}

func GetCourseDetails(c *gin.Context) {
	courseClass := c.Param("class")
	courseCategory := c.Param("category")
	courseID := c.Param("id")
	var err error

	conn := connections.FirebaseTutorialConnection()
	dsnap, err := conn.Collection("courses/" + courseClass + "/" + courseCategory).Doc(courseID).Get(context.Background())
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
