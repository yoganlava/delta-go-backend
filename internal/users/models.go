package users

type LoginStrategy string
type GenderType string

const (
	LOCAL    LoginStrategy = "local"
	FACEBOOK               = "facebook"
	TWITTER                = "twitter"
	GOOGLE                 = "google"
)

const (
	MALE   GenderType = "male"
	FEMALE            = "female"
	OTHER             = "other"
)

type UserEntity struct {
	Id            int8
	Username      string
	Email         string
	Password      string
	Verified      bool
	Social_id     string
	Strategy      LoginStrategy
	First_name    string
	Last_name     string
	Phone_number  string
	Gender        GenderType
	Refresh_token string
	Created_at    string
	Updated_at    string
}
