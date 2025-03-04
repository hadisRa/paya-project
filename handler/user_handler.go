package handler

import (
	"net/http"
	"paya/config"
	"paya/middleware"
	"paya/models"
	"paya/repository"
	"paya/service"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserSrv   service.User
	CacheRepo repository.CacheInterface
	UserRepo  repository.UserInterface
}

func NewUserHandler(srv service.User, cacheRepo repository.CacheInterface, userRepo repository.UserInterface) *UserHandler {
	return &UserHandler{
		UserSrv:   srv,
		CacheRepo: cacheRepo,
		UserRepo:  userRepo,
	}
}

func (h *UserHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			middleware.RequestCounterMiddleware(http.StatusBadRequest, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "@handler.user_handler.Register",
				"message": err.Error(),
			})
			return
		}

		otp := h.CacheRepo.GenerateRandomOTP()
		err := h.CacheRepo.Set(ctx, user.Username, otp, 5*time.Minute)
		if err != nil {
			middleware.RequestCounterMiddleware(http.StatusInternalServerError, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "@handler.user_handler.Register.GenerateRandomOTP",
				"message": "Failed to store OTP",
			})
			return
		}

		// TODO: Send the OTP to the user via email/SMS

		// Now, let's assume the user has submitted the OTP for verification
		var otpRequest struct {
			Username string `json:"username"`
			OTP      string `json:"otp"`
		}

		if err := ctx.ShouldBindJSON(&otpRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storedOTP, err := h.CacheRepo.Get(ctx, otpRequest.Username)
		if err == redis.Nil {
			middleware.RequestCounterMiddleware(http.StatusBadRequest, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "@handler.user_handler.Register.Get",
				"message": "OTP has expired or does not exist",
			})
			return
		} else if err != nil {
			middleware.RequestCounterMiddleware(http.StatusInternalServerError, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "@handler.user_handler.Register.Get",
				"message": "Failed to retrieve OTP",
			})
			return
		}

		if storedOTP != otpRequest.OTP {
			middleware.RequestCounterMiddleware(http.StatusBadRequest, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "@handler.user_handler.Register",
				"message": "Invalid OTP",
			})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			middleware.RequestCounterMiddleware(http.StatusInternalServerError, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "@handler.user_handler.Register.GenerateFromPassword",
				"message": "Failed to hash password",
			})
			return
		}
		user.Password = string(hashedPassword)

		if err := h.UserSrv.CreateUser(&user); err != nil {
			middleware.RequestCounterMiddleware(http.StatusInternalServerError, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "@handler.user_handler.Register.CreateUser",
				"message": "Failed to create user",
			})
			return
		}

		token, err := middleware.GenerateJWT(user.ID)
		if err != nil {
			middleware.RequestCounterMiddleware(http.StatusInternalServerError, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "@handler.user_handler.Register.CreateUser",
				"message": "Failed to generate token",
			})
			return
		}

		middleware.RequestCounterMiddleware(http.StatusOK, ctx.Request.Method, ctx.FullPath())
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"token":  token,
		})
	}
}

func (h *UserHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.User

		if err := ctx.ShouldBindJSON(&input); err != nil {
			middleware.RequestCounterMiddleware(http.StatusBadRequest, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   "@handler.user_handler.Login",
				"message": err.Error(),
			})
			return
		}

		user, err := h.UserSrv.GetUser(input.Username)
		if err != nil {
			middleware.RequestCounterMiddleware(http.StatusUnauthorized, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"error":   "@handler.user_handler.Login.GetUser",
				"message": "Invalid credentials",
			})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			middleware.RequestCounterMiddleware(http.StatusUnauthorized, ctx.Request.Method, ctx.FullPath())
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"error":   "@handler.user_handler.Login.CompareHashAndPassword",
				"message": "Invalid credentials",
			})
			return
		}

		token := generateToken(user.ID)

		middleware.RequestCounterMiddleware(http.StatusOK, ctx.Request.Method, ctx.FullPath())
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"token":  token,
		})
	}
}

func generateToken(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})
	tokenString, _ := token.SignedString([]byte(config.Cfg.JWT.Secret))
	return tokenString
}
