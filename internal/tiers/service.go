package tiers

import (
	"context"
	"main/internal/dto"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TierService struct {
	pool *pgxpool.Pool
}

type ITierService interface {
	CreateTier(t dto.CreateTierDTO) error
	FetchProjectTiers(id int) ([]*entity.Tier, error)
	FetchTier(id int) (entity.Tier, error)
}

func (ts TierService) CreateTier(t dto.CreateTierDTO) error {
	_, err := ts.pool.Exec(context.Background(), "insert into tier (title, description, cover_image_id, price, project_id, created_at, updated_at) values ($1,$2,$3,$4,$5,now(),now())", t.Name, t.Description, t.CoverImageID, t.Price, t.ProjectID)
	if err != nil {
		return err
	}
	return nil
}

func (ts TierService) FetchProjectTiers(id int) ([]*entity.Tier, error) {
	var t []*entity.Tier
	err := pgxscan.Select(context.Background(), ts.pool, &t, "select title, description, cover_image_id, price, project_id, created_at, updated_at from tier where project_id=$1", id)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (ts TierService) FetchTier(id int) (entity.Tier, error) {
	var t entity.Tier
	err := pgxscan.Get(context.Background(), ts.pool, &t, "select title, description, cover_image_id, price, project_id, created_at, updated_at from tier where id=$1", id)
	if err != nil {
		return entity.Tier{}, err
	}
	return t, nil
}
