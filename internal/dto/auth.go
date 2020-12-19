package dto

type AuthRegister struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthLogin struct {
	Credential string `json:"credential" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
type LoginPayload struct {
	JWT string `json:"jwt"`
}
