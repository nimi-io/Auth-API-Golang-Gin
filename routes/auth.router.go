package routes

import (
	"Auth-API/controllers"

	"github.com/gin-gonic/gin"
)


func AuthRoutes(incommingRoutes *gin.RouterGroup){
    authGroup := incommingRoutes.Group("/auth")

	go  authGroup.POST("/register", controllers.Signup)
	 authGroup.POST("/login", controllers.Login)

}