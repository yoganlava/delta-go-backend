package creators

import (
	"context"
	"errors"
	"fmt"
	"main/db"
	"main/internal/dto"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vmihailenco/msgpack/v5"
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

	val, err := db.Cache().Get(context.Background(), fmt.Sprintf(`/creators/%v`, id)).Result()
	err = msgpack.Unmarshal([]byte(val), &c)
	if c.ID > 0 {
		if c.UserID != user_id {
			c.CreatorRank = entity.CreatorRank{}
		}
		return c, nil
	}
	err = pgxscan.Get(context.Background(), cs.pool, &c, `
	 select c.id,c.name,c.avatar_image_id,
	 c.user_id,c.updated_at,c.created_at, 
	 JSON_BUILD_OBJECT('bio',cp.bio) as creator_profile,
	 JSON_BUILD_OBJECT('name',cr.name,'fee',cr.fee) as creator_rank
	 JSON_AGG(JSON_BUILD_OBJECT('name',p.name,'page_url',p.page_url,'id',p.id,'avatar',f.location,'banner',banner.location)) as projects
	 from creator c
	 inner join creator_rank cr on cr.id = c.creator_rank_id
	 inner join creator_profile cp on cp.creator_id = c.id
	 inner join project p on p.creator_id = c.id
	 inner join file f on f.id = p.avatar_image_id 
	 inner join file banner on banner.id = p.banner_image_id
	 where c.id = $1
	 group by c.id,cp.creator_id,cr.id`, id)
	if err != nil {
		return entity.Creator{}, err
	}

	if c.UserID != user_id {
		c.CreatorRank = entity.CreatorRank{}
	}
	marsh, err := msgpack.Marshal(&c)

	go db.Cache().Set(context.Background(), fmt.Sprintf(`/creators/%v`, id), marsh, 0)
	return c, nil
}

func (cs CreatorService) SearchCreators(name string, limit int, offset int) ([]*entity.SearchCreator, error) {
	var creators []*entity.SearchCreator
	err := pgxscan.Select(context.Background(), cs.pool, &creators, `
	select
	c.id, c.name, c.avatar_image_id, f.location as avatar
	from creator c
	where
	name like '%$1%'
	inner join file f on f.id = c.avatar_image_id
	limit $2 offset $3
	`,
		name, limit, offset,
	)
	return creators, err
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
	return creator, nil
}

func (cs CreatorService) UpdateCreator(c dto.UpdateCreatorDTO) (entity.Creator, error) {
	var creator entity.Creator

	trans, err := cs.pool.Begin(context.Background())

	if err != nil {
		trans.Rollback(context.Background())
		return entity.Creator{}, err
	}
	var userID int
	trans.QueryRow(context.Background(), `select user_id from creator where id = $1`, c.ID).Scan(&userID)
	if userID != c.UserID {
		trans.Rollback(context.Background())
		return entity.Creator{}, errors.New("このクリエーターを更新する権限がありません")
	}
	_, err = trans.Exec(context.Background(), `
	update creator
	set name = $1,bio = $2, avatar_image_id=$3,creator_rank_id=$4
	where id = $5 and user_id = $6
	`, c.Name, c.Bio, c.AvatarImageID, c.CreatorRankID, c.ID, c.UserID)

	if err != nil {
		trans.Rollback(context.Background())
		return entity.Creator{}, err
	}
	trans.Commit(context.Background())
	if err != nil {
		fmt.Print(err.Error())
		return entity.Creator{}, err
	}
	return creator, nil
}
