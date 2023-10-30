package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthenticateRoute (c *gin.Context) {

	c.Next()
}