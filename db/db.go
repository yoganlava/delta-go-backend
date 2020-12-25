package db

import (
	"context"
	"fmt"
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
	poolConfig.MinConns = 0
	poolConfig.MaxConns = int32(maxConn)
	poolConfig.MaxConnLifetime = time.Hour

	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	return db
}
