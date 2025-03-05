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

	ticketFilter := bson.D{{"_id", ticketID}}
	var Ticket structs.Ticket
	err = ticketCollection.FindOne(context.TODO(), ticketFilter).Decode(&Ticket)
	if err != nil {
		log.Println("Ошибка при поиске тикета:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера при поиске тикета"})
		return
	}

	userFilter := bson.D{{"name", Username}}
	var user structs.Customer
	err = userCollection.FindOne(context.TODO(), userFilter).Decode(&user)
	if err != nil {
		log.Println("Ошибка при поиске пользователя:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера при поиске пользователя"})
		return
	}

	update := bson.D{{"$push", bson.D{{"AssignedTicket", Ticket}}}}
	result, err := userCollection.UpdateOne(context.TODO(), bson.D{{"_id", user.ID}}, update)
	if err != nil {
		log.Println("Ошибка при обновлении пользователя:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера при обновлении пользователя"})
		return
	}

	deleteFilter := bson.D{{"_id", ticketID}}
	deleter_result, err := ticketCollection.DeleteOne(context.TODO(), deleteFilter)
	if err != nil {
		log.Println("Ошибка при удалении тикета:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера при удалении тикета"})
		return
	}

	if deleter_result.DeletedCount == 0 {
		log.Println("Тикет не найден для удаления")
		c.JSON(http.StatusNotFound, gin.H{"error": "Тикет не найден"})
		return
	}

	log.Println("Это результат обновления: ", result)
	log.Println("Тикет добавлен пользователю")
	c.JSON(http.StatusOK, gin.H{"message": "Тикет успешно добавлен пользователю"})
}
