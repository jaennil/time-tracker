package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/config"
)

func NewPostgres(config *config.Config) (*pgx.Conn, error) {
	connUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName, config.DBSSLMode)
	conn, err := pgx.Connect(context.Background(), connUrl)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
