package database

import (
	"fmt"
	"log"

	"github.com/enohr/rinha-backend-2024-q1/internal/infra/config"
	"github.com/jackc/pgx"
)

func NewDatabase(config config.Database) *pgx.Conn {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Db,
	)

	pgxConfig, err := pgx.ParseConnectionString(dbUrl)

	if err != nil {
		log.Fatal(err)
	}

	db, err := pgx.Connect(pgxConfig)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
