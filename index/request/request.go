package request

import (
	"context"
	"log"
	"main/index/structs"
	"net/http"
	"sync"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection
var initOnce sync.Once
var ticketCollection *mongo.Collection

func init() {
	initOnce.Do(func() {
		initiaseMongoDB()
	})
}

func initiaseMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return
	}

	log.Println("Connected to MongoDB!")

	userCollection = client.Database("Task").Collection("user")
	ticketCollection = client.Database("Task").Collection("ticket")
}

func RegistrationRequest(c *gin.Context) {
	name := c.PostForm("name_user")
	password := c.PostForm("password_user")
	log.Println("Данные :", name, password)
	if name == "" || password == "" {
		c.JSON(500, gin.H{
			"error": "Все поля должны быть заполнены",
		})
		return
	}
	var existingUser structs.Customer
	filter := bson.D{{"name", name}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if existingUser.Name == name {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Такой пользователь уже есть в базе данных",
		})
		return
	}
	users := structs.Customer{
		Name:     name,
		Password: password,
		Type_:    "Пользователь",
	}
	result, err := userCollection.InsertOne(context.TODO(), users)
	if err != nil {
		log.Panic(err)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("Ошибка приведения InsertedID к ObjectID")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервера",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", users.ID.Hex())
	session.Set("username", users.Name)
	session.Set("type", users.Type_)
	err = session.Save()
	log.Printf("Inserted document with ID: %s", oid.Hex())
	c.JSON(http.StatusOK, gin.H{
		"message": "Пользователь успешно добавлен",
	})

}

func RequestEnter(c *gin.Context) {
	name := c.PostForm("name_user")
	password := c.PostForm("password_user")
	if name == "" || password == "" {
		c.JSON(500, gin.H{
			"error": "Заполните все поля",
		})
	}
	var Users structs.Customer
	filter := bson.D{{"name", name}, {"password", password}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&Users)
	if err == mongo.ErrNoDocuments {
		c.JSON(200, gin.H{
			"message": "Пользователь не найден",
		})
	}
	session := sessions.Default(c)
	session.Set("user_id", Users.ID.Hex())
	session.Set("username", Users.Name)
	session.Set("type", Users.Type_)
	err = session.Save()
	if err != nil {
		log.Printf("Ошибка при сохранении сессии: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"eror": "Ошибка сервера",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Вход успешен",
	})
}

func RequestTicket() []structs.Ticket { // Добавил возврат ошибки
	var Tickets []structs.Ticket

	filter := bson.D{{"status", "open"}}

	cursor, err := ticketCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Ошибка при выполнении Find():", err)
		return nil
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var ticket structs.Ticket
		err := cursor.Decode(&ticket)
		if err != nil {
			log.Println("Ошибка при декодировании документа:", err)
			return nil
		}
		Tickets = append(Tickets, ticket)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Ошибка после итерации по курсору:", err)
		return nil
	}
	return Tickets
}
