package posts

import (
	"context"
	"errors"
	"main/db"
	"main/internal/dto"
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

func (ps PostService) DeleteProjectPost(deleteProjectPostDTO dto.DeleteProjectPostDTO) error {
	var creatorID int
	err := ps.pool.QueryRow(context.Background(), `
	select
	id as creatorID
	from creator
	where user_id = $1
	`,
		deleteProjectPostDTO.UserID,
	).Scan(&creatorID)
	if err != nil {
		return errors.New("全く許可しません")
	}
	_, err = ps.pool.Exec(context.Background(), `
	delete from post
	where id = $1
	`,
		creatorID,
	)
	return err
}
