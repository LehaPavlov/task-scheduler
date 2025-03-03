package request

import (
	"context"
	"log"
	"main/index/structs"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTicket(c *gin.Context) {
	session := sessions.Default(c)
	nameInterface := session.Get("username")
	name := nameInterface.(string)
	description := c.PostForm("description")
	title := c.PostForm("title")
	if description == "" || title == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Заполните все поля!",
		})
	}
	ticket := structs.Ticket{
		ID:          primitive.NewObjectID(),
		Creator:     name,
		Created:     time.Now(),
		Description: description,
		Title:       title,
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
