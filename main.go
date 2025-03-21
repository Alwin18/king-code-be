package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Alwin18/king-code/config"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Application panicked: %v", r)
		}
	}()

	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	db := config.NewDatabase(cfg)
	app := config.NewGin(cfg)
	config.MigrateDatabase(db)

	// setup bootstrap
	config.Bootstrap(&config.BootstrapConfig{
		DB:  db.Debug(),
		App: app,
		Cfg: cfg,
	})

	// create signal handling to gracefully shut down the application
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// start the server in a goroutine
	go func() {
		err := app.Run(":" + cfg.ServerPort)
		if err != nil {
			log.Printf("Error running server: %v", err)
		}
	}()

	// wait for shutdown signal
	<-sigs
	log.Println("Shutting down...")

	// Close the database connection
	config.CloseDB(db)

	log.Println("Application exited gracefully.")
}
