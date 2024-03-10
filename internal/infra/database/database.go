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
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Db,
	)

	cfg, err := pgxpool.ParseConfig(dbUrl)

	db, err := pgxpool.NewWithConfig(context.Background(), cfg)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	return db
}
