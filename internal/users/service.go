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
	// Register registers the user and returns jwt token
	// Login registers the user and returns jwt token
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

// Register user
// func (us UserService) Register(user RegisterRequest) gin.H {
// 	_, err := us.pool.Exec(context.Background(), "insert into users (username, password) VALUES ($1, $2)", user.Username, user.Password)
// 	if err != nil {
// 		return gin.H{
// 			"error": err.Error(),
// 		}
// 	}
// 	return gin.H{
// 		"message": "User created",
// 	}
// }

//Login user
func (us UserService) FetchSelf(id int) (entity.SelfUser, error) {
	var u entity.SelfUser
	// err:= us.pool.QueryRow(context.Background(),).Scan(&u.Id,&u.Password,&u.Username,&u.)

	err := pgxscan.Get(context.Background(), us.pool, &u, "select * from users where username = $1 or email = $1", id)
	if err != nil {
		return entity.SelfUser{}, err
	}
	return u, nil
}
