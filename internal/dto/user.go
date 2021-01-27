package dto

type ResetPasswordDTO struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}
