package middleware

import (
	"log"
	"net/http"

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
