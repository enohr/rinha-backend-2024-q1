package api

import (
	"github.com/gofiber/fiber/v3"
)

type Router struct {
	App         *fiber.App
	Handlers    *ClientesHandlers
	Middlewares *ClientesMiddleware
}

func NewRouter(app *fiber.App, handlers *ClientesHandlers, middlewares *ClientesMiddleware) *Router {
	return &Router{
		App:         app,
		Handlers:    handlers,
		Middlewares: middlewares,
	}
}

func (r *Router) Start() {
	group := r.App.Group("clientes/:id", r.Middlewares.ValidateGroup)
	group.Post("/transacoes", r.Handlers.HandleTransacoes, r.Middlewares.ValidateTransacoes)
	group.Get("/extrato", r.Handlers.HandleExtrato)
}
