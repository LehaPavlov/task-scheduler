package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson: "name`
	Password       string             `bson: "password`
	Type_          string             `bson: "type_`
	AssignedTicket []Ticket           `bson:"AssignedTicketIDs,omitempty"`
}

type Ticket struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Creator     string             `bson: "creator"`
	Created     time.Time          `bson: "created"`
	Title       string             `bson: "title"`
	Description string             `bson: "description"`
	Status      string             `bson: "status"`
}
type Comment struct {
	UserID    primitive.ObjectID `bson:"userID"`
	Text      string             `bson:"text"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type Chat struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Participants []primitive.ObjectID `bson:"participants"`
	Messages     []Message            `bson:"messages"`
}

type Message struct {
	Sender    primitive.ObjectID `bson:"senderID"`
	Text      string             `bson:"text"`
	CreatedAt time.Time          `bson:"createdAt"`
}
