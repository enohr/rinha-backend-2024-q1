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
	group := r.App.Group("clientes/:id")
	group.Post("/transacoes", r.Handlers.HandleTransacoes)
	group.Get("/extrato", r.Handlers.HandleExtrato)
}
