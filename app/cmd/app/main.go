package main

import (
	"go-bolvanka/internal/app"
	"go-bolvanka/internal/config"
	"log"
)

func main() {
	log.Print("config initializing")
	cfg, err := config.NewConfig()

	log.Print("logger initializing")

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	a, _ := app.NewApp(cfg)

	a.Run()
}
