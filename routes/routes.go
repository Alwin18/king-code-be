package routes

import (
	"github.com/Alwin18/king-code/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
	levelHandler *handlers.LevelHandler,
	progressHandler *handlers.ProgressHandler,
	challengeHandler *handlers.ChallengeHandler,
	wsHandler gin.HandlerFunc,
) *gin.Engine {
	router := gin.Default()

	// User Routes
	user := router.Group("/users")
	{
		user.POST("/register", userHandler.RegisterUser)
		user.GET("/:id", userHandler.GetUserByID)
	}

	// Level Routes
	level := router.Group("/levels")
	{
		level.GET("/", levelHandler.GetAllLevels)
	}

	// Progress Routes
	progress := router.Group("/progress")
	{
		progress.GET("/:userID", progressHandler.GetUserProgress)
	}

	// Challenge Routes
	challenge := router.Group("/challenges")
	{
		challenge.GET("/level/:levelID", challengeHandler.GetChallengesByLevel)
	}

	// WebSocket Route (Real-time coding & multiplayer)
	router.GET("/ws", wsHandler)

	return router
}
