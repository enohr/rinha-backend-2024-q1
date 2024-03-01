package database

import (
	"context"
	"time"

	"github.com/enohr/rinha-backend-2024-q1/internal/domain/clientes"
	"github.com/enohr/rinha-backend-2024-q1/internal/infra/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(config config.Database) *Repository {
	db := NewDatabase(config)

	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveTransacao(ctx context.Context, id string, t *clientes.Transacao) (*clientes.Saldo, error) {
	var saldo clientes.Saldo
	rows, err := r.db.Query(ctx, transacaoQuery, id, t.Valor, t.Tipo, t.Descricao)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&saldo.Total, &saldo.Limite); err != nil {
			return nil, err
		}
	}

	return &saldo, nil
}

func (r *Repository) GetExtrato(ctx context.Context, id string) (*clientes.Extrato, error) {
	var extrato clientes.Extrato
	var transacao clientes.Transacao
	extrato.Transacoes = make([]clientes.Transacao, 0)
	extrato.Saldo.DataExtrato = time.Now()

	rows, err := r.db.Query(ctx, extratoQuery, id)

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
