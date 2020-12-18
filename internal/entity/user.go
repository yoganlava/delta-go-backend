package entity

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
	ID           int8   `json:"id"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string
	Verified     bool
	SocialID     string
	Strategy     LoginStrategy
	FirstName    string
	LastName     string
	PhoneNumber  string
	Gender       GenderType
	RefreshToken string
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string
}

type SafeUser struct {
	ID        int8       `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Gender    GenderType `json:"gender"`
	CreatedAt string     `json:"created_at"`
}

type SelfUser struct {
	ID          int8   `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	FirstName   string
	LastName    string
	PhoneNumber string
	Gender      GenderType
	Strategy    LoginStrategy
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string
	Verified    bool
}
