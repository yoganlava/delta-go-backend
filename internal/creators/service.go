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
	CreateCreator(c dto.CreateCreatorDTO) (entity.Creator, error)
}

func New() CreatorService {
	return CreatorService{db.Connection()}
}

type CreatorService struct {
	pool *pgxpool.Pool
}

func (cs CreatorService) FetchCreator(id int, user_id int) (entity.Creator, error) {
	var c entity.Creator
	err := pgxscan.Get(context.Background(), cs.pool, &c, `
	 select c.id,c.name,c.avatar_image_id,
	 c.user_id,c.updated_at,c.created_at, 
	 JSON_BUILD_OBJECT('bio',cp.bio) as creator_profile,
	 JSON_BUILD_OBJECT('name',cr.name,'fee',cr.fee) as creator_rank
	 from creator c
	 inner join creator_rank cr on cr.id = c.creator_rank_id
	 inner join creator_profile cp on cp.creator_id = c.id
	 where c.id = $1
	 group by c.id,cp.creator_id,cr.id`, id)
	if err != nil {
		return entity.Creator{}, err
	}

	if c.UserID != user_id {
		c.CreatorRank = entity.CreatorRank{}
	}
	return c, nil
}

func (cs CreatorService) CreateCreator(c dto.CreateCreatorDTO) (entity.Creator, error) {
	var creator entity.Creator
	err := pgxscan.Get(context.Background(), cs.pool, &creator, "insert into creator (name,avatar_image_id,user_id,creator_rank_id,created_at,updated_at) values($1,$2,$3,1,now(),now()) returning id,name,avatar_image_id,user_id,creator_rank_id,updated_at,created_at", c.Name, c.AvatarImageID, c.UserID)
	go func() {
		_, err := cs.pool.Exec(context.Background(), "insert into creator_profile (bio,creator_id)  values($1,$2)", c.Bio, creator.ID)
		if err != nil {
			fmt.Println(err.Error())
		}

	}()
	if err != nil {
		fmt.Print(err.Error())
		return entity.Creator{}, err
	}
	fmt.Println("Ended")
	return creator, nil
}
