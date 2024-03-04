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
	type TransacaoResponse struct {
		Limite int `json:"limite"`
		Saldo  int `json:"saldo"`
	}

	t := c.Locals("transaction").(*clientes.Transacao)
	id := c.Params("id")

	saldo, err := ch.service.SaveTransacao(c.Context(), id, t)

	switch err {
	case clientes.ErrLimiteInsuficiente:
		return c.Status(422).SendString("Insufficient limit")
	default:
		return c.Status(500).SendString("An error occurred")
	case nil:
		break
	}

	transacao := TransacaoResponse{
		Limite: saldo.Limite,
		Saldo:  saldo.Total,
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
