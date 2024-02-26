package api

import (
	"github.com/enohr/rinha-backend-2024-q1/internal/domain/clientes"
	"github.com/gofiber/fiber/v3"
)

type Router struct {
	App      *fiber.App
	Handlers *clientes.ClientesHandlers
}

func NewRouter(app *fiber.App, handlers *clientes.ClientesHandlers) *Router {
	return &Router{
		App:      app,
		Handlers: handlers,
	}
}

func (r *Router) Start() {
	r.App.Post("/clientes/:id/transacoes", r.Handlers.HandleTransacoes)
	r.App.Get("/clientes/:id/extrato", r.Handlers.HandleExtrato)
}
