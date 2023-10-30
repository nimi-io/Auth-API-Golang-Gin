package routes

import (
	"Auth-API/controller"

	"github.com/gin-gonic/gin"
)

func  AuthRouts(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.POST("/register", controller.Register )
	router.POST("/login", controller.Login) 
}