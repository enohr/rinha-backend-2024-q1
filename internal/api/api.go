package api

import (
	"github.com/enohr/rinha-backend-2024-q1/internal/domain/clientes"
	"github.com/enohr/rinha-backend-2024-q1/internal/infra/config"
	"github.com/enohr/rinha-backend-2024-q1/internal/infra/database"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func NewAPI(cfg *config.Config) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	repository := database.NewRepository(cfg.Database)
	service := clientes.NewClientesService(repository)
	handlers := NewClientesHandlers(service)
	middlewares := NewClientesMiddleware()

	r := NewRouter(app, handlers, middlewares)
	r.Start()

	return app
}
