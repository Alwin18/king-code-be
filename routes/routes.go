package routes

import (
	"github.com/Alwin18/king-code/handlers"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App                *gin.Engine
	CORSMiddleware     gin.HandlerFunc
	RecoveryMiddleware gin.HandlerFunc
	UserHandler        *handlers.UserHandler
	LevelHandler       *handlers.LevelHandler
	ProgressHandler    *handlers.ProgressHandler
	ChallengeHandler   *handlers.ChallengeHandler
	WsHandler          gin.HandlerFunc
}

func (r *RouteConfig) Setup() {
	v1 := r.App.Group("api/v1")
	r.App.Use(r.CORSMiddleware, r.RecoveryMiddleware)

	// User Routes
	user := v1.Group("users")
	{
		user.POST("/register", r.UserHandler.RegisterUser)
		user.GET("/:id", r.UserHandler.GetUserByID)
	}

	// Level Routes
	level := v1.Group("/levels")
	{
		level.GET("/", r.LevelHandler.GetAllLevels)
	}

	// Progress Routes
	progress := v1.Group("/progress")
	{
		progress.GET("/:userID", r.ProgressHandler.GetUserProgress)
	}

	// Challenge Routes
	challenge := v1.Group("/challenges")
	{
		challenge.GET("/level/:levelID", r.ChallengeHandler.GetChallengesByLevel)
	}

	// WebSocket Route (Real-time coding & multiplayer)
	v1.GET("/ws", r.WsHandler)
}
