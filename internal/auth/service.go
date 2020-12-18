package auth

import (
	"context"
	"errors"
	"main/db"
	"main/internal/dto"
	"main/internal/entity"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IAuthService interface {
	// Register registers the user and returns jwt token
	Register(user dto.AuthRegister) string
	// Login registers the user and returns jwt token
	Login(user dto.AuthLogin) string
	CreateToken(id int) string
	VerifyToken(tokenString string) (int error)
}
type AuthService struct {
	pool *pgxpool.Pool
}

type UserClaim struct {
	jwt.MapClaims
	id  int
	exp int64
}

func New() AuthService {
	return AuthService{db.Connection()}
}

// CreateToken for user
func CreateToken(id int) string {
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"id":  id,
	// 	"exp": time.Now().Add(time.Minute * 30).Unix(),
	// })
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		id:  id,
		exp: time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// VerifyToken and return user id or error
func VerifyToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Something went wrong when parsing")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return -1, err
	}
	if !token.Valid {
		return -1, errors.New("Token expired")
	}
	return token.Claims.(*UserClaim).id, nil
}

func (auth AuthService) Register(request dto.AuthRegister) error {
	_, err := auth.pool.Exec(context.Background(), "insert into users (username, password) VALUES ($1, $2)", request.Username, request.Password)

	if err != nil {
		return err
	}

	return nil
}

//Login user
func (auth AuthService) Login(request dto.AuthLogin) (entity.SafeUser, error) {
	var u entity.SafeUser
	// err:= us.pool.QueryRow(context.Background(),).Scan(&u.Id,&u.Password,&u.Username,&u.)
	err := pgxscan.Get(context.Background(), auth.pool, &u, "select * from users where username = $1 or email = $1", u.Username)
	if u.ID == 0 {
		return u, nil
	}
	if err != nil {
		return u, err
	}
	return u, nil

}
