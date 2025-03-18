package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Alwin18/king-code/utils"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If the request method is OPTIONS, respond and return without calling the next handler
		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(http.StatusOK)
			return
		}

		c.Next()
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.String(http.StatusInternalServerError, "Internal Server Error")
				log.Printf("Panic: %s", err)
			}
		}()
		c.Next()
	}
}

// AuthMiddleware is used to protect routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Pastikan token dalam format "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Parse token
		claims, err := utils.ParseToken(tokenParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":         "Invalid or expired token",
				"refresh_token": "Please use /auth/refresh-token",
			})
			c.Abort()
			return
		}

		// Simpan user_id di context agar bisa digunakan di handler
		c.Set("user_id", claims.UserID)

		c.Next() // Lanjut ke handler berikutnya
	}
}
