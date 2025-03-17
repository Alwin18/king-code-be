package main

import (
	"log"

	"github.com/Alwin18/king-code/config"
)

func main() {
	config.LoadEnv()
	config.InitDatabase()
	config.MigrateDatabase()
	log.Println("database migrated successfully!")
}
