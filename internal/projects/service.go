package projects

import (
	"context"
	"errors"
	"main/db"
	"main/internal/dto"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProjectService struct {
	pool *pgxpool.Pool
}

// New create new Project Service
func New() ProjectService {
	return ProjectService{db.Connection()}
}

type IProjectService interface {
	CreateProject(p dto.CreateProjectDTO) (int, error)
	FetchProject(id int) (entity.Project, error)
}

func (ps ProjectService) FetchProject(id int) (entity.Project, error) {
	var p entity.Project
	err := pgxscan.Get(context.Background(), ps.pool, &p, "select id, name, page_url, description, creating, banner_image_id, cover_image_id, creator_id, cover_id, category_id, setting, created_at, updated_at")
	if err != nil {
		return entity.Project{}, err
	}
	return p, nil
}

func (ps ProjectService) CreateProject(p dto.CreateProjectDTO) error {
	_, err := ps.pool.Exec(context.Background(), "insert into project (name, page_url, description, creating, creator_id, cover_id, category_id, setting, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, now(), now())", p.Name, p.PageURL, p.Description, p.Creating, p.CreatorID, p.CoverID, p.CategoryID)
	if err != nil {
		return err
	}
	return nil
}

//! Temporary, does not belong here
func (ps ProjectService) GetCreatorIDFromUserID(userID int) (int, error) {
	creatorID := 0
	err := ps.pool.QueryRow(context.Background(), "select id from creator where user_id=$1", userID).Scan(&creatorID)
	if err != nil {
		return -1, err
	}
	if creatorID <= 0 {
		return -1, errors.New("No creator found")
	}
	return creatorID, nil
}
