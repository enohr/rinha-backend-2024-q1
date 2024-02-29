package api

import (
	"github.com/enohr/rinha-backend-2024-q1/internal/domain/clientes"
	"github.com/gofiber/fiber/v3"
)

type ClientesHandlers struct {
	service *clientes.ClientesService
}

func NewClientesHandlers(service *clientes.ClientesService) *ClientesHandlers {
	return &ClientesHandlers{
		service: service,
	}
}

func (ch *ClientesHandlers) HandleTransacoes(c fiber.Ctx) error {
	t := new(clientes.Transacao)
	if err := c.Bind().Body(t); err != nil {
		return err
	}

	id := c.Params("id")
	// TODO: Verify if the ID exists

	transacao, err := ch.service.SaveTransacao(c.Context(), id, t)

	if err != nil {
		if err == clientes.ErrLimiteInsuficiente {
			return c.Status(422).SendString("Limite insuficiente")
		}
	}

	return c.JSON(transacao)
}

func (ch *ClientesHandlers) HandleExtrato(c fiber.Ctx) error {
	id := c.Params("id")

	extrato, err := ch.service.GetExtrato(c.Context(), id)
	if err != nil {
		return c.SendString("Error")
	}

	return c.JSON(extrato)
}
