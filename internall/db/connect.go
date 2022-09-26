package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

func Connect(dsn string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.WithError(err).Info("can't connection to db")
	}

	return conn
}
