package data

import (
	"context"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/luccasbarros/the-service/cmd/api/config"
)

func InitPool() (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Env.User, config.Env.Password, config.Env.Host, config.Env.Port, config.Env.Name,
	)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	poolConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &pgtype.UUID{},
			Name:  "uuid",
			OID:   pgtype.UUIDOID,
		})
		return nil
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
