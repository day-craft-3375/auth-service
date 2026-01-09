package main

import (
	"log"

	"github.com/day-craft-3375/auth-service/config"
	"github.com/day-craft-3375/auth-service/internal/app"
)

func main() {
	// Config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Run
	app.Run(cfg)
}
