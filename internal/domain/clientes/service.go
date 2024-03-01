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

func (s *ClientesService) SaveTransacao(ctx context.Context, id string, t *Transacao) (*Saldo, error) {
	if t.Tipo == "d" {
		t.Valor = -t.Valor
	}

	return s.repository.SaveTransacao(ctx, id, t)
}

func (s *ClientesService) GetExtrato(ctx context.Context, id string) (*Extrato, error) {
	return s.repository.GetExtrato(ctx, id)
}
