package middleware

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func AuthenticateRoute(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort()
		return
	}

	tokenString = tokenString[7:]

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort()
		return
	}

	secretKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the token's signing method and return the secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}



	// if token.Valid {
	// 	token.
	// 	c.Next()
	// } else {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	// 	c.Abort()
	// 	return
	// }


	   if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Access the token's content here
        username := claims["username"].(string)
        userID := claims["userId"].(any)
		ID := claims["ID"].(any)
		Firstname := claims["Firstname"].(string)
		Lastname := claims["Lastname"].(string)
		Email := claims["Email"].(string)
		Phone := claims["Phone"].(string)

        // You can access other claims as needed

        // Now, you can use username and userID in your route handling logic
        c.Set("username", username)
        c.Set("userId", userID)
		c.Set("ID", ID)
		c.Set("Firstname", Firstname)
		c.Set("Lastname", Lastname)
		c.Set("Email", Email)
		c.Set("Phone", Phone)
        c.Next()
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        c.Abort()
    }
	return
}
