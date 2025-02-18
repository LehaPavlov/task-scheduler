package render

import (
	"log"
	"main/index/request"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RenderMainPage(c *gin.Context) {
	session := sessions.Default(c)
	usernameInterface := session.Get("username")
	typeUserInterface := session.Get("type")
	isLoggedInInterface, exists := c.Get("isLoggedIn")

	isLoggedIn := false
	if exists {
		value, ok := isLoggedInInterface.(bool)
		if !ok {
			log.Println("Ошибка при смене типа isLoggedIn")
			isLoggedIn = false
		} else {
			isLoggedIn = value
		}
	}

	username := "Гость"
	if usernameInterface != nil {
		usernameValue, ok := usernameInterface.(string)
		if ok {
			username = usernameValue
		} else {
			log.Println("Ошибка при смене типа username")
		}
	}

	type_ := ""
	if typeUserInterface != nil {
		typeValue, ok := typeUserInterface.(string)
		if ok {
			type_ = typeValue
		} else {
			log.Println("Ошибка при смене типа type")
		}
	}

	tickets := request.RequestTicket()

	data := gin.H{
		"isLogged": isLoggedIn,
		"Name":     username,
		"Type":     type_,
	}
	log.Println("Это type", type_)
	if type_ == "Администратор" {
		data["Ticket"] = tickets
	}

	c.HTML(http.StatusOK, "main.html", data)
}

func RenderRegistration(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", nil)
}

func RenderEnterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "enter.html", nil)
}

func RenderCreateTicket(c *gin.Context) {
	c.HTML(http.StatusOK, "ticket.html", nil)
}
