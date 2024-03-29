package api

import (
	"github.com/enohr/rinha-backend-2024-q1/internal/domain/clientes"
	"github.com/enohr/rinha-backend-2024-q1/internal/infra/config"
	"github.com/enohr/rinha-backend-2024-q1/internal/infra/database"
	"github.com/gofiber/fiber/v3"
)

func NewAPI(cfg *config.Config) *fiber.App {
	app := fiber.New()

	repository := database.NewRepository(cfg.Database)
	service := clientes.NewClientesService(repository)
	handlers := NewClientesHandlers(service)

	r := NewRouter(app, handlers)
	r.Start()

	return app
}
