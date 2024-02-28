package database

import (
	"context"

	"github.com/enohr/rinha-backend-2024-q1/internal/domain/clientes"
	"github.com/enohr/rinha-backend-2024-q1/internal/infra/config"
	"github.com/jackc/pgx"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(config config.Database) *Repository {
	db := NewDatabase(config)

	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveTransacao(ctx context.Context, id string, t *clientes.Transacao) error {
	return nil
}

func (r *Repository) GetExtrato(ctx context.Context, id string) (*clientes.Extrato, error) {

	return nil, nil
}
