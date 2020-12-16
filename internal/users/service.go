package users

import (
	"context"
	"main/db"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

// IUserService encapsulates logic
type IUserService interface {
	// Register registers the user and returns jwt token
	Register(user RegisterRequest) string
	// Login registers the user and returns jwt token
	Login(user LoginRequest) string
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
func (us UserService) Register(request RegisterRequest) gin.H {
	_, err := us.pool.Exec(context.Background(), "insert into users (username, password) VALUES ($1, $2)", request.Username, request.Password)
	if err != nil {
		return gin.H{
			"error": err.Error(),
		}
	}
	return gin.H{
		"message": "User created",
	}
}

//Login user
func (us UserService) Login(user LoginRequest) gin.H {
	return gin.H{
		"message": "User logged",
	}
}
