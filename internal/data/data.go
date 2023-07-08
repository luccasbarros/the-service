package data

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Data struct {
	db *pgxpool.Pool
}

var qb = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func New(pool *pgxpool.Pool) *Data {
	return &Data{pool}
}
