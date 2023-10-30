package controllers

import (
	"Auth-API/helpers"
	"Auth-API/models"
	"Auth-API/utils"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "golang.org/x/telemetry/counter"
)

func Signup(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request", //"user": user,
		})
		return
	}

	validateErr := validate.Struct(user)
	if validateErr != nil {
		c.JSON(400, gin.H{
			"message": validateErr.Error(),
		})
		return

	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil || count > 0 {
		log.Println("email.userCollection.CountDocuments", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	defer cancel()
	if err != nil || count > 0 {
		log.Println("phone.userCollection.CountDocuments", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if count > 0 {
		log.Println("Email already exists", count)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exists",
		})
		return
	}
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	objectID := primitive.NewObjectID()
	user.ID = objectID

	user.User_id = objectID.Hex()

	token, err := helpers.GenerateJWTToken(user)
	if err != nil {
		log.Println("helpers.GenerateAccessToken", err)
		
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	user.Token = token

	user.Password = utils.HashedPassword(user.Password)

	result, err := userCollection.InsertOne(ctx, user)

	if err != nil {
		log.Println("userCollection.InsertOne", err)
		
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, gin.H{
		"message": "User Created", "result": result},
	)

}
func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login",
	})

}
