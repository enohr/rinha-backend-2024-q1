package clientes

import "github.com/enohr/rinha-backend-2024-q1/internal/infra/database"

type ClientesService struct {
}

func NewClientesService(repository *database.Repository) *ClientesService {
	return &ClientesService{}
}
