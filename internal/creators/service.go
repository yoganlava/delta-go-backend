package creators

import (
	"context"
	"fmt"
	"main/db"
	"main/internal/dto"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ICreatorService interface {
	FetchCreator(id int) (entity.Creator, error)
}

func New() CreatorService {
	return CreatorService{db.Connection()}
}

type CreatorService struct {
	pool *pgxpool.Pool
}

func (cs CreatorService) FetchCreator(id int) (entity.Creator, error) {
	var c entity.Creator

	err := pgxscan.Get(context.Background(), cs.pool, &c, "select id,name, bio,banner_image_id,cover_image_id,user_id,updated_at,created_at from creators where id = $1", id)
	if err != nil {
		return entity.Creator{}, err
	}
	return c, nil
}

func (cs CreatorService) CreateCreator(c dto.CreateCreatorDTO) (entity.Creator, error) {
	var creator entity.Creator
	err := pgxscan.Get(context.Background(), cs.pool, &creator, "insert into creator (name, bio,avatar_image_id,user_id,creator_rank_id,created_at,updated_at) values($1,$2,$3,$4,1,now(),now())", c.Name, c.Bio, c.AvatarImageID, c.UserID)
	fmt.Errorf("%s", err.Error())
	if err != nil {
		return entity.Creator{}, err
	}

	return creator, nil
}
