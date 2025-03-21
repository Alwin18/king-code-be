package config

import (
	"github.com/Alwin18/king-code/handlers"
	"github.com/Alwin18/king-code/pkg/middleware"
	"github.com/Alwin18/king-code/repositories"
	"github.com/Alwin18/king-code/routes"
	"github.com/Alwin18/king-code/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB  *gorm.DB
	App *gin.Engine
	Cfg *Config
}

func Bootstrap(cfg *BootstrapConfig) {
	// Init Repository
	userRepo := repositories.NewUserRepository(cfg.DB)
	levelRepo := repositories.NewLevelRepository(cfg.DB)
	progressRepo := repositories.NewProgressRepository(cfg.DB)
	challengeRepo := repositories.NewChallengeRepository(cfg.DB)
	tokenRepo := repositories.NewTokenRepository(cfg.DB)
	leadBoardRepo := repositories.NewLeaderboardRepository(cfg.DB)
	submissionRepo := repositories.NewSubmissionRepository(cfg.DB)

	// Init Service
	userService := services.NewUserService(userRepo)
	levelService := services.NewLevelService(levelRepo)
	progressService := services.NewProgressService(progressRepo, levelRepo)
	challengeService := services.NewChallengeService(challengeRepo)
	leadboardService := services.NewLeaderboardService(leadBoardRepo)
	submissionService := services.NewSubmissionService(submissionRepo, challengeRepo)

	// Init Handler
	userHandler := handlers.NewUserHandler(userService)
	levelHandler := handlers.NewLevelHandler(levelService)
	progressHandler := handlers.NewProgressHandler(progressService)
	challengeHandler := handlers.NewChallengeHandler(challengeService)
	authHandler := handlers.NewAuthHandler(userService, tokenRepo)
	leadboardHandler := handlers.NewLeaderboardHandler(leadboardService)
	submissionHandler := handlers.NewSubmissionHandler(submissionService)

	// init websocket
	ws := handlers.WebSocketHandler

	// Setup Router
	routeConfig := routes.RouteConfig{
		App:                cfg.App,
		CORSMiddleware:     middleware.CORSMiddleware(),
		AuthMiddleware:     middleware.AuthMiddleware(),
		RecoveryMiddleware: middleware.RecoveryMiddleware(),
		UserHandler:        userHandler,
		LevelHandler:       levelHandler,
		AuthHandler:        authHandler,
		SubbmissionHandler: submissionHandler,
		ProgressHandler:    progressHandler,
		LeadboardHandler:   leadboardHandler,

		ChallengeHandler: challengeHandler,
		WsHandler:        ws,
	}

	routeConfig.Setup()
}
