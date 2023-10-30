package routes

import (
	"Auth-API/controller"

	"github.com/gin-gonic/gin"
)

type LocalAuthController struct {
    controller.AuthController
}

func (ac *LocalAuthController) RegisterUserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.POST("/register", ac.Register)
	router.POST("/login", ac.Login) 
}