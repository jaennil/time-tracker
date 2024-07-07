package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/pkg/database"
)

func NewPostgres(config database.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), dsn(config))
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func dsn(config database.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User, config.Password, config.Host, config.Port, config.Name, config.SSLMode)
}
