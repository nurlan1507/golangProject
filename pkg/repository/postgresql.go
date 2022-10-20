package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	DbName   string
	User     string
	Password string
}

func OpenDb(cfg *Config) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf("postgres://postgres:%+v@localhost:5432/%+v", cfg.Password, cfg.DbName)
	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
