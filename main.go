package main

import (
	"main/index/render"
	"main/index/request"
	"main/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	secret_key := []byte("secret_key")
	store := cookie.NewStore(secret_key)
	router := gin.Default()
	router.Use(sessions.Sessions("mysession", store))
	router.LoadHTMLGlob("templates/*")
	router.GET("/", middleware.AuthMiddleware(), render.MainPage)
	router.POST("/", request.Render_Ticket)
	router.GET("/registration", render.Registration)
	router.GET("/enter", render.EnterPage)
	router.POST("/registration", request.Registration)
	router.POST("/enter", request.Enter)
	router.GET("/create_ticket", render.CreateTicket)
	router.POST("/create_ticket", request.CreateTicket)
	router.Run()
}
