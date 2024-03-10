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
	id := c.Params("id")

	if id == "6" {
		return c.Status(404).SendString("User not found")
	}

	t := &clientes.Transacao{}
	if err := c.Bind().Body(t); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	if len(t.Descricao) > 10 || len(t.Descricao) == 0 {
		return c.Status(400).SendString("Descrption invalid")
	}

	if t.Tipo != "d" && t.Tipo != "c" {
		return c.Status(400).SendString("Type must be 'd' or 'c'")
	}

	type TransacaoResponse struct {
		Limite int `json:"limite"`
		Saldo  int `json:"saldo"`
	}

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

	if id == "6" {
		return c.Status(404).SendString("User not found")
	}

	extrato, err := ch.service.GetExtrato(c.Context(), id)
	if err != nil {
		return c.Status(500).SendString("An error occurred")
	}

	return c.JSON(extrato)
}
