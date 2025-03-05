package render

import (
	"log"
	"main/index/request"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
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

	tickets := request.Ticket()
	assigned_tickets := request.Assigned_tickets(c)
	data := gin.H{
		"isLogged": isLoggedIn,
		"Name":     username,
		"Type":     type_,
		"Ticket":   tickets,
	}
	log.Println("Это assigned tickets", assigned_tickets)
	if type_ == "Администратор" {
		data["AssignedTickets"] = assigned_tickets.AssignedTicket
	}

	c.HTML(http.StatusOK, "main.html", data)
}

func Registration(c *gin.Context) {
	c.HTML(http.StatusOK, "registration.html", nil)
}

func EnterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "enter.html", nil)
}

func CreateTicket(c *gin.Context) {
	c.HTML(http.StatusOK, "ticket.html", nil)
}
