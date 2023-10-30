package helpers

import (
	"Auth-API/models"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUID (c *gin.Context, userId string)(err error){	
	userType := c.GetString("user_type")
	uid:=  c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	} 

	err = CheckUserType(c, userType)
	if err != nil {
		return err
	}
return err
}

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil

	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}



func GenerateJWTToken(d models.User) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "user": d,
  }) 

  tokenString, err := token.SignedString([]byte("secret"))
  if err != nil {
    return "", err
  }

  return tokenString, nil
}