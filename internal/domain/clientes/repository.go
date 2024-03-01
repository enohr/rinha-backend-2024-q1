package clientes

import "context"

type Repository interface {
	GetExtrato(ctx context.Context, id string) (*Extrato, error)
	SaveTransacao(ctx context.Context, id string, t *Transacao) (*Saldo, error)
}
