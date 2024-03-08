package main

import (
	"fmt"

	"github.com/enohr/rinha-backend-2024-q1/internal/api"
	"github.com/enohr/rinha-backend-2024-q1/internal/infra/config"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Warn("Error loading .env file")
	}

	config := config.GetConfig()
	app := api.NewAPI(config)

	addr := fmt.Sprintf(":%s", config.API.Port)
	app.Listen(addr)
}
