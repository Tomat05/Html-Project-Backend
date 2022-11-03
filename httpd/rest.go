package rest

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"html-project-backend/database"
)

func Gin(collection *mongo.Collection) {
	r := gin.Default()

	r.GET("/users", database.FindUser(collection))
	r.POST("/create", database.CreateUser(collection))

	r.Run("localhost:8080")
}
