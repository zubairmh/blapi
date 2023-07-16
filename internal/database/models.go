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
	Author       string             `bson:"author" json:"author"`
	AuthorMetaID string             `bson:"meta_id" json:"meta_id"`
	Title        string             `bson:"title" json:"title"`
	Description  string             `bson:"description" json:"description"`
	Deadline     time.Time          `bson:"deadline" json:"deadline"`
	Bounty       float32            `bson:"bounty" json:"bounty"`
}

type Application struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Author       string             `bson:"author" json:"author"`
	AuthorMetaID string             `bson:"meta_id" json:"meta_id"`
	PostID       primitive.ObjectID `bson:"post_id" json:"post_id"`
	Description  string             `bson:"description" json:"description"`
}
