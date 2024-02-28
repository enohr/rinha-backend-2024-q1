package api

import (
	"github.com/gofiber/fiber/v3"
)

type Router struct {
	App      *fiber.App
	Handlers *ClientesHandlers
}

func NewRouter(app *fiber.App, handlers *ClientesHandlers) *Router {
	return &Router{
		App:      app,
		Handlers: handlers,
	}
}

func (r *Router) Start() {
	// TODO: Use Group
	r.App.Post("/clientes/:id/transacoes", r.Handlers.HandleTransacoes)
	r.App.Get("/clientes/:id/extrato", r.Handlers.HandleExtrato)
}
