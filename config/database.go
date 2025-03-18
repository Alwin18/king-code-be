package config

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *Config) *gorm.DB {
	if cfg.SSLMode == "" {
		cfg.SSLMode = "prefer"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err = connection.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected and pinged database successfully")

	maxIdleConns := 30
	if cfg.SetMaxIdleConns != "" {
		maxIdleConns, _ = strconv.Atoi(cfg.SetMaxIdleConns)
	}

	MaxOpenConns := 100
	if cfg.SetMaxOpenConns != "" {
		MaxOpenConns, _ = strconv.Atoi(cfg.SetMaxOpenConns)
	}

	maxLifeTimeConnection := 300
	if cfg.SetMaxLifeTime != "" {
		maxLifeTimeConnection, _ = strconv.Atoi(cfg.SetMaxLifeTime)
	}

	connection.SetMaxIdleConns(maxIdleConns)
	connection.SetMaxOpenConns(MaxOpenConns)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal("Failed to close connection to database:", err)
	}
	dbSQL.Close()
}
