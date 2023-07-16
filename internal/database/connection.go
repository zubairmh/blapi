package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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

func GetUser(meta_id string) (User, error) {
	user := User{}
	err := users.FindOne(ctx, bson.M{meta_id: meta_id}).Decode(&user)
	log.Printf("Get User %s", meta_id)
	return user, err
}

func CreateUser(meta_id string, name string) User {
	u, err := GetUser(meta_id)
	if err != nil {
		log.Printf("User %s exists", u.MetaMaskID)
		return u
	}
	u.MetaMaskID = meta_id
	u.Name = name
	u.ID = primitive.NewObjectID()
	users.InsertOne(ctx, u)
	log.Printf("User %s created", u.MetaMaskID)
	return u
}

func CreateBounty(meta_id string, bounty float32, title string, description string, deadline time.Time) Post {
	u, _ := GetUser(meta_id)
	post := Post{
		AuthorMetaID: meta_id,
		Author:       u.Name,
		Title:        title,
		Description:  description,
		Bounty:       bounty,
		Deadline:     deadline,
		ID:           primitive.NewObjectID(),
	}
	posts.InsertOne(ctx, post)
	return post
}

func GetBounties() []Post {
	cur, err := posts.Find(ctx, bson.D{})
	var results []Post
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	log.Println("Request all posts")
	return results
}

func GetApplications(post_id primitive.ObjectID) []Application {
	cur, err := applications.Find(ctx, bson.D{primitive.E{Key: "post_id", Value: post_id}})
	var results []Application
	if err = cur.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	log.Printf("Applications searched for %s", post_id)
	return results
}

func CreateApplication(meta_id string, description string, post_id primitive.ObjectID) {
	u, _ := GetUser(meta_id)
	applications.InsertOne(ctx,
		Application{
			Author:       u.Name,
			AuthorMetaID: meta_id,
			Description:  description,
			PostID:       post_id,
		})
}
