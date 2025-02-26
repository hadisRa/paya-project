package middleware

import (
	"net/http"
	"strconv"
	"time"

	"paya/metrics"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}

func RequestDurationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start).Seconds()
		metrics.RequestDuration.WithLabelValues(c.Request.Response.Status, c.Request.Method, c.FullPath()).Observe(duration)
	}
}

func RequestCounterMiddleware(status int, method, path string) {
	strStatus := strconv.Itoa(status)
	metrics.TaskCreationCounter.WithLabelValues(strStatus, method, path).Inc()
}
