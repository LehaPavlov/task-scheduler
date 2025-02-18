package structs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson: "name`
	Password string             `bson: "password`
	Type_    string             `bson: "type`
}

type Ticket struct {
	Creator     string    `bson: "creator"`
	Created     time.Time `bson: "created"`
	Description string    `bson: "description"`
	Status      string    `bson: "status"`
}
