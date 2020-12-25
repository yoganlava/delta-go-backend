package files

import "github.com/jackc/pgx/v4/pgxpool"

type FileService struct {
	pool *pgxpool.Pool
}

type IFileService interface{}
