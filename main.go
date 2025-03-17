package main

import (
	"log"

	"github.com/Alwin18/king-code/config"
	"github.com/Alwin18/king-code/handlers"
	"github.com/Alwin18/king-code/repositories"
	"github.com/Alwin18/king-code/routes"
	"github.com/Alwin18/king-code/services"
)

func main() {
	config.LoadEnv()
	config.InitDatabase()
	// config.MigrateDatabase()

	// Init Repository
	userRepo := repositories.NewUserRepository(config.DB)
	levelRepo := repositories.NewLevelRepository(config.DB)
	progressRepo := repositories.NewProgressRepository(config.DB)
	challengeRepo := repositories.NewChallengeRepository(config.DB)

	// Init Service
	userService := services.NewUserService(userRepo)
	levelService := services.NewLevelService(levelRepo)
	progressService := services.NewProgressService(progressRepo)
	challengeService := services.NewChallengeService(challengeRepo)

	// Init Handler
	userHandler := handlers.NewUserHandler(userService)
	levelHandler := handlers.NewLevelHandler(levelService)
	progressHandler := handlers.NewProgressHandler(progressService)
	challengeHandler := handlers.NewChallengeHandler(challengeService)

	// Setup Router
	router := routes.SetupRouter(userHandler, levelHandler, progressHandler, challengeHandler, handlers.WebSocketHandler)

	log.Println("Server is running on port 8080...")
	router.Run(":8080")
}
