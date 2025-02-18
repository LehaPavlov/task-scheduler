package request

import (
	"context"
	"log"
	"main/index/structs"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateTicket(c *gin.Context) {
	session := sessions.Default(c)
	nameInterface := session.Get("username")
	name := nameInterface.(string)
	description := c.PostForm("description")
	if description == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Форма пуста!",
		})
	}
	ticket := structs.Ticket{
		Creator:     name,
		Created:     time.Now(),
		Description: description,
		Status:      "open",
	}
	err, result := ticketCollection.InsertOne(context.TODO(), ticket)
	if err != nil {
		log.Println("Ошибка в запросе: ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
