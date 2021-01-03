package auth

import (
	"context"
	"errors"
	"main/db"
	"main/internal/dto"
	"main/internal/email"
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
	// CreateToken(id int) string
	// VerifyToken(tokenString string) (int error)
}
type AuthService struct {
	pool *pgxpool.Pool
}

type UserClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func New() AuthService {
	return AuthService{db.Connection()}
}

// CreateToken for user
func CreateToken(id int) dto.CreateTokenDTO {
	now := time.Now()
	now.Add(time.Hour * 24)
	var claim = UserClaim{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.UnixNano(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}
	return dto.CreateTokenDTO{
		JWT: tokenString,
		EXP: now.UnixNano(),
	}
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
		//Expired token
		return -1, errors.New("トークンの期限が切れています")
	}
	return token.Claims.(*UserClaim).ID, nil
}

// Register user
func (auth AuthService) Register(request *dto.AuthRegister) (entity.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	var user entity.User
	pgxscan.Get(context.Background(), auth.pool, &user, `
		select * from users where lower(email) = $1 or lower(username) = $2
	`, request.Email, request.Username)

	if user.ID > 0 {
		return entity.User{}, errors.New("ユーザー名またはEメールがもう使われています")
	}
	err = pgxscan.Get(context.Background(), auth.pool, &user, "insert into users (email,username, password,verified,created_at,updated_at,strategy,avatar) VALUES ($1, $2,$3,$4,now(),now(),'local','https://img.jpmtl.com/default_profile.png') returning id,username,password,verified,created_at,avatar",
		request.Email, request.Username, string(hashed), false)

	if err != nil {
		return entity.User{}, err
	}
	token := CreateToken(user.ID)
	email.SendVerificationEmail(request.Email, request.Username, token.JWT)
	return user, nil
}

//Login user
func (auth AuthService) Login(request dto.AuthLogin) (entity.AuthUser, error) {
	var u = entity.AuthUser{}
	// err:= us.pool.QueryRow(context.Background(),).Scan(&u.Id,&u.Password,&u.Username,&u.)
	err := pgxscan.Get(context.Background(), auth.pool, &u, `
	select 
	u.id,u.username,u.password,u.email,u.created_at,u.gender,u.verified,
	f.location as avatar
	from users u
	inner join file f on f.id = u.avatar_image_id
	where u.username = $1 or u.email = $1`, request.Credential)
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

// Rudimentary
// func (auth AuthService) sendVerificationEmail(user_id int) error {
// 	var u = entity.AuthUser{}
// 	err := pgxscan.Get(context.Background(), auth.pool, &u, `
// 	select
// 	u.id,u.username,u.password,u.email,u.created_at,u.gender,u.verified
// 	from users u
// 	where u.id = $1`, user_id)
// 	if err != nil {
// 		return err
// 	}
// 	token := CreateToken(u.ID)

// }
