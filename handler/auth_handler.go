package handler

import (
	"net/http"
	"paya/config"
	"paya/models"
	"paya/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service service.User
}

func NewUserHandler(srv service.User) *UserHandler {
	return &UserHandler{
		service: srv,
	}
}
func (h *UserHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//TODO creteUser service and GetUser service
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(hashedPassword)

		// database.DB.Create(&user)

		ctx.JSON(http.StatusOK, user)
	}
}
func (h *UserHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		var input models.User
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		// 	return
		// }
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		token := generateToken(user.ID)
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}
func generateToken(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})
	tokenString, _ := token.SignedString([]byte(config.Cfg.JWT.Secret))
	return tokenString
}
