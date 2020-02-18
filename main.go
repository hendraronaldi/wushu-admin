package main

import (
	"encoding/gob"
	"time"
	"work/wushu-backend/chatbot/line"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gob.Register(map[string]interface{}{})
	gob.Register(time.Time{})
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(CORSMiddleware())
	line.LoadAllContext()
	SetupRouter(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
