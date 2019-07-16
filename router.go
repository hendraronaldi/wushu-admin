package main

import (
	"fmt"
	"net/http"
	"work/wushu-backend/modules/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, X-Requested-With, Access-Control-Request-Method, Access-Control-Request-Headers, Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cache-Control")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func SetupRouter(router *gin.Engine) {
	Router(router)
}

func Router(router *gin.Engine) {
	UsersRouter(router)
	AdminRouter(router)
}

func UsersRouter(router *gin.Engine) {
	router.PUT("/edituser", controller.EditUser)
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	// router.POST("/welcome", controller.Welcome)
	// router.POST("/logout", controller.Logout)
	// router.POST("/userprofile", controller.GetUser)
	// router.POST("/users", controller.PostUser)
	// router.PUT("/users", controller.UpdateUser)
	// router.DELETE("/users", controller.DeleteUser)
}

func AdminRouter(router *gin.Engine) {
	router.POST("/validate", controller.ValidateUser)
	router.POST("/reject", controller.RejectUser)
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// You'd normally redirect to login page
			c.JSON(http.StatusBadRequest, gin.H{"response": "Invalid session token"})
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}