package creators

import (
	"context"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ICreatorService interface {
	FetchCreator(id int) (entity.Creator, error)
}

type CreatorService struct {
	pool *pgxpool.Pool
}

func (cs CreatorService) FetchCreator(id int) (entity.Creator, error) {
	var c entity.Creator

	err := pgxscan.Get(context.Background(), cs.pool, &c, "select * from creators where id=$1", id)
	if err != nil {
		return entity.Creator{}, err
	}
	return c, nil
}
