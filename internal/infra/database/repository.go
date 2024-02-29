package database

import (
	"context"
	"time"

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
	var extrato clientes.Extrato
	var transacao clientes.Transacao
	extrato.Transacoes = make([]clientes.Transacao, 0)
	extrato.Saldo.DataExtrato = time.Now()

	rows, err := r.db.QueryEx(ctx, extratoQuery, nil, id)
	if err != nil {
		return nil, err
	}

	// TODO: Check if user exists
	for rows.Next() {
		if err := rows.Scan(&extrato.Saldo.Total, &extrato.Saldo.Limite, &transacao.Valor, &transacao.Tipo, &transacao.Descricao, &transacao.RealizadaEm); err != nil {
			return nil, err
		}
		extrato.Transacoes = append(extrato.Transacoes, transacao)
	}
	return &extrato, nil
}
