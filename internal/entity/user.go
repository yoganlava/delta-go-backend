package entity

import "time"

//LoginStrategy  enum
type LoginStrategy string

//GenderType enum
type GenderType string

//login strategy enum definition
const (
	LOCAL    LoginStrategy = "local"
	FACEBOOK               = "facebook"
	TWITTER                = "twitter"
	GOOGLE                 = "google"
)

//gender type enum definition

const (
	//GenderType enum
	MALE   GenderType = "male"
	FEMALE            = "female"
	OTHER             = "other"
)

//User model
type User struct {
	ID            int           `json:"id"`
	Email         string        `json:"email"`
	Username      string        `json:"username"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	Avatar        string        `json:"avatar"`
	PhoneNumber   string        `json:"phone_number"`
	Gender        GenderType    `json:"gender"`
	Strategy      LoginStrategy `json:"strategy"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Verified      bool          `json:"verified"`
	SocialID      string        `json:"social_id"`
	Password      string        `json:"password"`
	AvatarImageID string        `json:"avatar_image_id"`
}

type AuthUser struct {
	ID          int           `json:"id"`
	Email       string        `json:"email" `
	Username    string        `json:"username"`
	Password    string        `json:"password"`
	Verified    bool          `json:"verified"`
	Gender      GenderType    `json:"gender"`
	CreatedAt   time.Time     `json:"created_at"`
	Strategy    LoginStrategy `json:"strategy"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	PhoneNumber string        `json:"phone_number"`
	Avatar      string        `json:"avatar"`
}

type SelfUser struct {
	ID          int           `json:"id"`
	Email       string        `json:"email"`
	Username    string        `json:"username"`
	FirstName   *string       `json:"first_name"`
	LastName    *string       `json:"last_name"`
	Avatar      string        `json:"avatar"`
	PhoneNumber *string       `json:"phone_number"`
	Gender      GenderType    `json:"gender"`
	Strategy    LoginStrategy `json:"strategy"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Verified    bool          `json:"verified"`
}
