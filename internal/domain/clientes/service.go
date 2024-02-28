package clientes

import (
	"context"
)

type ClientesService struct {
	repository Repository
}

func NewClientesService(repository Repository) *ClientesService {
	return &ClientesService{
		repository: repository,
	}
}

func (s *ClientesService) SaveTransacao(ctx context.Context, id string, t *Transacao) error {
	return s.repository.SaveTransacao(ctx, id, t)
}
