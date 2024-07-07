package main

import (
	"log"

	"github.com/jaennil/time-tracker/config"
	"github.com/jaennil/time-tracker/internal/app"
)

func main() {

	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("fatal error while creating config: %s", err)
	}

	app.Run(config)
}
