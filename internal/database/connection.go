package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var posts *mongo.Collection
var applications *mongo.Collection
var users *mongo.Collection

var ctx context.Context = context.TODO()

func init() {
	godotenv.Load()
	uri := os.Getenv("MONGO")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Loaded")
	users = client.Database("bountydb").Collection("users")
	applications = client.Database("bountydb").Collection("applications")
	posts = client.Database("bountydb").Collection("posts")
}

func CreateApplication(meta_id string, description string, post_id primitive.ObjectID) {
	applications.InsertOne(ctx,
		Application{})
}
