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
	router.GET("/", middleware.AuthMiddleware(), render.RenderMainPage)
	router.GET("/registration", render.RenderRegistration)
	router.GET("/enter", render.RenderEnterPage)
	router.POST("/registration", request.RegistrationRequest)
	router.POST("/enter", request.RequestEnter)
	router.GET("/create_ticket", render.RenderCreateTicket)
	router.POST("/create_ticket", request.CreateTicket)
	router.Run()
}
