package database

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html-project-backend/httpd/models"
	"log"
	"net/http"
)

// Connect to mongodb database cos SQL is for nerds
func Connect() (*mongo.Client, *mongo.Collection) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// Connect to database
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("html-project").Collection("users")
	fmt.Println(collection)

	return client, collection
}

// Check if user exists
func userExists(collection *mongo.Collection, user models.User) (bool, *models.User) {
	result := new(models.User)
	filter := bson.D{{"name", user.Name}}

	collection.FindOne(context.TODO(), filter).Decode(&result)
	return result.Name != "", result
}

// Insert data into database
func CreateUser(collection *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var Input models.User
		if err := ctx.ShouldBindJSON(&Input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		exists, user := userExists(collection, Input)

		if exists {
			ctx.JSON(http.StatusOK, gin.H{"data": "User exists!"})
			ctx.JSON(http.StatusOK, gin.H{"data": user})
			return
		}

		insertResult, err := collection.InsertOne(context.TODO(), Input)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"data": insertResult})
		fmt.Println(insertResult)
	}
}

// Update stuff in database
func UpdateUser(collection *mongo.Collection) {
	filter := bson.D{{"name", "Tom"}}
	update := bson.D{
		{"$set", bson.D{
			{"email", "different@email.com"},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func FindUser(collection *mongo.Collection) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var result models.User
		filter := bson.D{{"name", "Tom"}}

		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusOK, gin.H{"data": "User Doesn't exist"})
			return
		}
		fmt.Printf("Found a single document: %+v\n", result)
		ctx.JSON(http.StatusOK, gin.H{"data": result})
	}
}

//func main() {
//	client, collection := connect()
//	//Insert(collection)
//	//Update(collection)
//	Find(collection)
//
//	if err := client.Disconnect(context.TODO()); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("connection closed")
//}
