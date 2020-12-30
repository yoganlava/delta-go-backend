package tiers

import (
	"context"
	"fmt"
	"main/db"
	"main/internal/dto"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vmihailenco/msgpack/v5"
)

type TierService struct {
	pool *pgxpool.Pool
}

type ITierService interface {
	CreateTier(t dto.CreateTierDTO) error
	FetchProjectTiers(id int) ([]*entity.Tier, error)
	FetchTier(id int) (entity.Tier, error)
}

func New() TierService {
	return TierService{db.Connection()}
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
	val, err := db.Cache().Get(context.Background(), fmt.Sprintf(`/tiers/project/%v`, id)).Result()
	if err == redis.Nil {
		err = pgxscan.Select(context.Background(), ts.pool, &t, "select title, description, cover_image_id, price, project_id, created_at, updated_at from tier where project_id=$1", id)
		if err != nil {
			return nil, err
		}
		marsh, err := msgpack.Marshal(&t)
		if err != nil {
			return nil, err
		}
		go db.Cache().Set(context.Background(), fmt.Sprintf(`/tiers/project/%v`, id), marsh, 0)
		return t, nil
	} else if err != nil {
		return nil, err
	}
	err = msgpack.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (ts TierService) FetchTier(id int) (entity.Tier, error) {
	var t entity.Tier

	val, err := db.Cache().Get(context.Background(), fmt.Sprintf(`/tiers/%v`, id)).Result()
	if err != nil {
		return entity.Tier{}, err
	}
	err = msgpack.Unmarshal([]byte(val), &t)
	if err != nil {
		return entity.Tier{}, err
	}
	if t.ID > 0 {
		return t, nil
	}

	err = pgxscan.Get(context.Background(), ts.pool, &t, "select title, description, cover_image_id, price, project_id, created_at, updated_at from tier where id=$1", id)
	if err != nil {
		return entity.Tier{}, err
	}
	marsh, err := msgpack.Marshal(&t)

	go db.Cache().Set(context.Background(), fmt.Sprintf(`/tiers/%v`, id), marsh, 0)
	return t, nil
}
