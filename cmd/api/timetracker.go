package main

import (
	"log"

	"github.com/jaennil/time-tracker/config"
	_ "github.com/jaennil/time-tracker/docs"
	"github.com/jaennil/time-tracker/internal/app"
)

//	@title		Time Tracker API
//	@version	1.0

//	@host		localhost:8081
//	@BasePath	/v1

//	@accept		json
//	@produce	json

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("fatal error while creating config: %s", err)
	}

	app.Run(cfg)
}
