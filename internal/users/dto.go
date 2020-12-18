package users

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type SafeUserEntity struct {
	Id        int8
	Email     string
	Username  string
	CreatedAt string
	UpdatedAt string
}
