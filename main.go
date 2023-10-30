package main

import (
	"Auth-API/controller"
	"Auth-API/routes"
	"Auth-API/services"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	userservice services.UserService
	// authService services.
	// controller     controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoConnect   *mongo.Client
	err            error
	authController controller.AuthController
)

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx = context.TODO()

	MONGO_URL := os.Getenv("MONGO_URL")
	log.Println(MONGO_URL)
	mongoOptions := options.Client().ApplyURI(MONGO_URL)
	mongoConnect, err = mongo.Connect(ctx, mongoOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoConnect.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB connection Established")

	userCollection = mongoConnect.Database("go-rest-api").Collection("users")
    services.NewUserService(userCollection, ctx)
	server = gin.Default()

}
func main() {
	defer mongoConnect.Disconnect(ctx)

	basePath := server.Group("/api/v1")
	routes.AuthRouts(basePath)

	log.Fatal(server.Run(":3010"))

}
