package request

import (
	"context"
	"log"
	"main/index/structs"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Render_Ticket(c *gin.Context) {
	session := sessions.Default(c)
	Username := session.Get("username")
	log.Println("Поиск по сессии ", Username)
	ticket_ID := c.PostForm("ticket_id")
	log.Println(ticket_ID)

	ticketID, err := primitive.ObjectIDFromHex(ticket_ID)
	if err != nil {
		log.Println("Ошибка конвертации ID тикета:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID тикета"})
		return
	}

	var user structs.Customer
	filter := bson.D{{"name", Username}}
	err = userCollection.FindOne(context.TODO(), filter).Decode(&user)
	update := bson.D{{"$push", bson.D{{"assigned_ticket_ids", ticketID}}}}

	result, err := userCollection.UpdateOne(context.TODO(), bson.D{{"_id", user.ID}}, update)
	if err != nil {
		log.Println("Ошибка при обновлении пользователя:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера при обновлении пользователя"})
		return
	}
	log.Println("Это результат обновления: ", result)
	log.Println("Тикет добавлен пользователю")
	c.JSON(http.StatusOK, gin.H{"message": "Тикет успешно добавлен пользователю"})
}
