package controller

import (
	"Auth-API/models"
	"Auth-API/services"
	"Auth-API/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService services.UserService
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(userService services.UserService) AuthController {

	return AuthController{UserService: userService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var request LoginRequest
	ctx.ShouldBindJSON(&request)

	user, err := c.UserService.GetByEmail(request.Username)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comppare := utils.ComparePasswords(user.Password, request.Password)

	if !comppare == false {
		ctx.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var user models.UserModel
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := c.UserService.GetByEmail(user.Username)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Username Unavailable"})
		return
	}

	isStrongPassword := utils.IsStrongPassword(user.Password)

	if !isStrongPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password is not strong enough"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}
	user.Password = hashedPassword
	c.UserService.Create(user)
	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}
