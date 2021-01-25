package posts

import (
	"context"
	"main/db"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostService struct {
	pool *pgxpool.Pool
}

type IPostService interface {
	FetchCreatorPosts(id int) (entity.Post, error)
}

// New create new Post Service
func New() PostService {
	return PostService{db.Connection()}
}

func (ps PostService) FetchProjectPosts(creatorID int) ([]*entity.Post, error) {
	var p []*entity.Post
	err := pgxscan.Select(context.Background(), ps.pool, &p, `
	select
	id, title, content, mature, project_id, submit_id, created_at, updated_at
	where id=$1`,
		creatorID,
	)
	return p, err
}
