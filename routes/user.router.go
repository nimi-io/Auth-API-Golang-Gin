package routes

import (
	"Auth-API/controllers"
	"Auth-API/middleware"

	"github.com/gin-gonic/gin"
)


func UserRoutes(incommingRoutes *gin.RouterGroup){	
	userGroup:= incommingRoutes.Group("/user")
	incommingRoutes.Use(middleware.AuthenticateRoute)

	userGroup.GET("/users",controllers.GetUsers)
	userGroup.GET("/user/:id",controllers.GetUserById)
	
}