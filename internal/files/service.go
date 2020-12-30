package files

import (
	"context"
	"io"
	"main/db"
	"mime/multipart"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

type FileService struct {
	pool *pgxpool.Pool
}

type IFileService interface {
	SaveFile(file *multipart.FileHeader) error
}

func New() FileService {
	return FileService{db.Connection()}
}

func (fs FileService) SaveFile(file *multipart.FileHeader) error {
	// ! LOCATION WILL BE BUCKET URL?
	// ! ALSO NO METADATA BESIDES HEADER INFO FOR NOW
	fileContent, _ := file.Open()
	buf := make([]byte, 512)
	if _, err := io.ReadFull(fileContent, buf); err != nil {
		return err
	}
	_, err := fs.pool.Exec(context.Background(), "insert into file (location, size, file_name, mime_type, created_at, updated_at, user_id, meta)", "", file.Size, file.Filename, http.DetectContentType(buf))
	return err
}
