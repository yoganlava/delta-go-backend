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
	"golang.org/x/crypto/bcrypt"
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
func (auth AuthService) CreateToken(id int) string {
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
func (auth AuthService) VerifyToken(tokenString string) (int, error) {
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
		return -1, errors.New("トークンの期限が切れています")
	}
	return token.Claims.(*UserClaim).id, nil
}

func (auth AuthService) Register(request *dto.AuthRegister) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	_, err = auth.pool.Exec(context.Background(), "insert into users (email,username, password,verified,created_at,updated_at,strategy) VALUES ($1, $2,$3,$4,now(),now(),'local')",
		request.Email, request.Username, string(hashed), false)

	if err != nil {
		return err
	}

	return nil
}

//Login user
func (auth AuthService) Login(request dto.AuthLogin) (entity.AuthUser, error) {
	var u = entity.AuthUser{}
	// err:= us.pool.QueryRow(context.Background(),).Scan(&u.Id,&u.Password,&u.Username,&u.)
	err := pgxscan.Get(context.Background(), auth.pool, &u, "select id,username,password,email,created_at,gender,verified from users where username = $1 or email = $1", request.Credential)
	hashedError := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(request.Password))
	if u.ID == 0 {
		return entity.AuthUser{}, nil
	}
	if hashedError != nil {
		return entity.AuthUser{}, hashedError
	}
	if err != nil {
		return entity.AuthUser{}, err
	}
	u.Password = ""
	return u, nil

}
