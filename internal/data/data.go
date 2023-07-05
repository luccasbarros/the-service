package data

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Data struct {
	db *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Data {
	return &Data{pool}
}

