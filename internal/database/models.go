package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	MetaMaskID string             `bson:"meta_id" json:"meta_id"`
}

type Post struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Author       string
	AuthorMetaID string
	Title        string
	Description  string
	Deadline     time.Time
	Bounty       float32
}

type Application struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Author       string
	AuthorMetaID string
	Description  string
}
