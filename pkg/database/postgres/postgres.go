package postgres

import (
	"fmt"

	"github.com/jaennil/time-tracker/pkg/database"
	"github.com/jmoiron/sqlx"
)

func New(config database.Config) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Name, config.Password, config.SSLMode)
	db, err := sqlx.Connect("postgres", connectionString)
	return db, err

}
