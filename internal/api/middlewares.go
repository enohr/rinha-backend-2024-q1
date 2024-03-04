package api

import (
	"github.com/enohr/rinha-backend-2024-q1/internal/domain/clientes"
	"github.com/gofiber/fiber/v3"
)

type ClientesMiddleware struct{}

func NewClientesMiddleware() *ClientesMiddleware {
	return &ClientesMiddleware{}
}

func (m *ClientesMiddleware) ValidateGroup(c fiber.Ctx) error {
	id := c.Params("id")

	if id == "6" {
		return c.Status(404).SendString("User not found")
	}

	return c.Next()
}

func (m *ClientesMiddleware) ValidateTransacoes(c fiber.Ctx) error {
	t := &clientes.Transacao{}
	if err := c.Bind().Body(t); err != nil {
		return c.Status(400).SendString("Invalid body")
	}
	c.Locals("transaction", t)

	if len(t.Descricao) > 10 || len(t.Descricao) == 0 {
		return c.Status(400).SendString("Descrption invalid")
	}

	if t.Tipo != "d" && t.Tipo != "c" {
		return c.Status(400).SendString("Type must be 'd' or 'c'")
	}

	return c.Next()
}
