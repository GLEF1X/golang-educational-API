package database

import (
	"context"
	"github.com/GLEF1X/golang-educational-API/core/logging"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConnectionPool(connectionURI string) *pgxpool.Pool {
	pgxConnPool, err := pgxpool.Connect(context.Background(), connectionURI)
	if err != nil {
		logging.GetLogger().Fatal(err)
	}
	return pgxConnPool
}
