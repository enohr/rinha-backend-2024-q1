package database

import (
	"context"
	"fmt"
	"log"

	"github.com/enohr/rinha-backend-2024-q1/internal/infra/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(config config.Database) *pgxpool.Pool {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Db,
	)

	db, err := pgxpool.New(context.Background(), dbUrl)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
