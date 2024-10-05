package db

import (
	"context"
	"fmt"
	"net"

	"github.com/EugeneTsydenov/chesshub-server/infrastructure/env"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn *pgx.Conn
}

func New() (*DB, error) {
	conn, err := initDB()
	if err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

func initDB() (*pgx.Conn, error) {
	username := env.Get("DATABASE_USERNAME")
	password := env.Get("DATABASE_PASSWORD")
	host := env.Get("DATABASE_HOST")
	port := env.Get("DATABASE_PORT")
	name := env.Get("DATABASE_NAME")

	hostPort := net.JoinHostPort(host, port)

	url := fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, hostPort, name)

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database at %s:%s (DB: %s): %w", host, port, name, err)
	}

	return conn, nil
}
