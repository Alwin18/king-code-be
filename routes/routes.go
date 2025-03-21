package routes

import (
	"github.com/Alwin18/king-code/handlers"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App                *gin.Engine
	CORSMiddleware     gin.HandlerFunc
	RecoveryMiddleware gin.HandlerFunc
	AuthMiddleware     gin.HandlerFunc
	UserHandler        *handlers.UserHandler
	LevelHandler       *handlers.LevelHandler
	ProgressHandler    *handlers.ProgressHandler
	SubbmissionHandler *handlers.SubmissionHandler
	ChallengeHandler   *handlers.ChallengeHandler
	AuthHandler        *handlers.AuthHandler
	LeadboardHandler   *handlers.LeaderboardHandler
	WsHandler          gin.HandlerFunc
}

func (r *RouteConfig) Setup() {
	v1 := r.App.Group("api/v1")
	r.App.Use(r.CORSMiddleware, r.RecoveryMiddleware)

	// Public Routes
	v1.POST("/users/register", r.UserHandler.RegisterUser)
	v1.POST("/users/login", r.AuthHandler.Login)

	// Protected Routes
	auth := v1.Group("/")
	// auth.Use(r.AuthMiddleware)

	auth.GET("/users/:id", r.UserHandler.GetUserByID)
	auth.GET("/levels", r.LevelHandler.GetAllLevels)
	auth.POST("/levels", r.LevelHandler.CreateLevel)
	auth.GET("/progress/:userID", r.ProgressHandler.GetUserProgress)
	auth.POST("progress", r.ProgressHandler.CreateProgress)
	auth.PATCH("/progress", r.ProgressHandler.UpdateUserProgress)
	auth.GET("/leaderboard", r.LeadboardHandler.GetLeaderboard)
	auth.GET("challenges/level/:level_id", r.ChallengeHandler.GetChallengesByLevel)
	auth.GET("challenges/:id", r.ChallengeHandler.GetChallengeByID)
	auth.POST("submissions/submit", r.SubbmissionHandler.SubmitCode)

	// WebSocket Route (Real-time coding & multiplayer)
	v1.GET("/ws", r.WsHandler)
}
