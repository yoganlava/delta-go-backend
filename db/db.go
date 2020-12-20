package db

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Connection gets connection of postgresql database
func Connection() (db *pgxpool.Pool) {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	maxConn, err := strconv.ParseInt(os.Getenv("MAX_POOL"), 10, 32)
	poolConfig.MinConns = 1
	poolConfig.MaxConns = int32(maxConn)
	poolConfig.MaxConnLifetime = time.Hour

	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)

	if err != nil {
		panic(err)
	}

	return db
}
