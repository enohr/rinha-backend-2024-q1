package clientes

import "github.com/gofiber/fiber/v3"

type ClientesHandlers struct {
}

func NewClientesHandlers(service *ClientesService) *ClientesHandlers {
	return &ClientesHandlers{}
}

func (ch *ClientesHandlers) HandleTransacoes(c fiber.Ctx) error {
	return c.SendString("Hello from Transacoes")
}

func (ch *ClientesHandlers) HandleExtrato(c fiber.Ctx) error {
	return c.SendString("Hello from Extrato")
}
