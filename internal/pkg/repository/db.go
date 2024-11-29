package repository

import (
	"context"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Params struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
	SslMode  string
}

func NewPostgresPool(ctx context.Context, cfg Params) (*pgxpool.Pool, error) {
	dbURL := "postgres://" + cfg.Username + ":" + url.QueryEscape(cfg.Password) + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DbName + "?sslmode=" + cfg.SslMode
	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	err = dbPool.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
