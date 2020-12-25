package users

import (
	"context"
	"main/db"
	"main/internal/entity"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

// IUserService encapsulates logic
type IUserService interface {
	FetchSelf(id int) (entity.SelfUser, error)
}

//UserService contains connection
type UserService struct {
	pool *pgxpool.Pool
}

// New create new User Service
func New() UserService {
	return UserService{db.Connection()}
}

//FetchSelf user
func (us UserService) FetchSelf(id int) (entity.SelfUser, error) {
	var u entity.SelfUser

	err := pgxscan.Get(context.Background(), us.pool, &u, `select 
	id,first_name,last_name,phone_number,username,email,created_at,updated_at,gender,strategy,verified 
	from users 
	where id=$1`, id)
	if err != nil {
		return entity.SelfUser{}, err
	}
	return u, nil
}

func (us UserService) isEmailAvailable(email string) bool {
	var u entity.User
	pgxscan.Get(context.Background(), us.pool, &u, `select * from users where lower(email) = lower($1)`, email)
	return u.ID < 1
}

func (us UserService) isUsernameAvailable(username string) bool {
	var u entity.User
	pgxscan.Get(context.Background(), us.pool, &u, `select * from users where lower(username) = lower($1)`, username)
	return u.ID < 1
}
