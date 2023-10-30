package controllers

import (
	"Auth-API/database"
	"Auth-API/helpers"
	"Auth-API/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.starlark.net/lib/time"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "Users")
var validate = validator.New()

func GetUsers(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var users []models.User

	cursor, err := userCollection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	defer cancel()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)

	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"data": users})

}

func GetUserById(c *gin.Context) {
	userId := c.Param("userId")

	helpers.MatchUserTypeToUID(c, userId)

	if err := helpers.MatchUserTypeToUID(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	defer cancel()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
