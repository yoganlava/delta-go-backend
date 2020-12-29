package dto

type AuthRegister struct {
	Email                string `json:"email" binding:"required,max=500,email"`
	Username             string `json:"username" binding:"required,min=2,max=30"`
	Password             string `json:"password" binding:"required,min=8,max=255"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,max=300"`
}

type AuthLogin struct {
	Credential string `json:"credential" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
type CreateTokenDTO struct {
	JWT string `json:"jwt"`
	EXP int64  `json:"exp_at"`
}

// type AuthVerificationDTO struct{
// 	Username string `json:`
// 	Email string
// }
