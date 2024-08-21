package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func GetConnection(cfg Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password,
		cfg.Database)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseConnection(db *sqlx.DB) error {
	return db.Close()
}
