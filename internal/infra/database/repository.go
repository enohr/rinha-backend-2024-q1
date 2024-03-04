package database

import (
	"context"
	"time"

	"github.com/enohr/rinha-backend-2024-q1/internal/domain/clientes"
	"github.com/enohr/rinha-backend-2024-q1/internal/infra/config"
	"github.com/jackc/pgx/v5/pgconn"
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
	transactionValue := t.Valor
	if t.Tipo == "d" {
		transactionValue = -transactionValue
	}

	rows := r.db.QueryRow(ctx, transacaoQuery, id, t.Valor, t.Tipo, t.Descricao, transactionValue)

	if err := rows.Scan(&saldo.Total, &saldo.Limite); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// "Constraint check violation"
			if pgErr.Code == "23514" {
				return nil, clientes.ErrLimiteInsuficiente
			}
		}
		return nil, err
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

	for rows.Next() {
		if err := rows.Scan(&extrato.Saldo.Total, &extrato.Saldo.Limite, &transacao.Valor, &transacao.Tipo, &transacao.Descricao, &transacao.RealizadaEm); err != nil {
			// this is a tricky for performance. this error means the user has no transactions
			// but we still need to return the info about the user.
			// Ideally, it should handle the Null values, but for performance, i'll do it this way
			if err.Error() == "can't scan into dest[2]: cannot scan NULL into *int" {
				return &extrato, nil
			}
			return nil, err
		}
		extrato.Transacoes = append(extrato.Transacoes, transacao)
	}
	return &extrato, nil
}
