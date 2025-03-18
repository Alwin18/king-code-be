package config

import "github.com/gin-gonic/gin"

func NewGin(cfg *Config) *gin.Engine {
	return gin.Default()
}
