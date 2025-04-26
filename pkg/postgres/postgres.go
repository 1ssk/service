package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func New(config Config) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w", err)

	}

	return conn, nil
}
