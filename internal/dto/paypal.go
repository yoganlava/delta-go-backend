package dto

type CreatePaypalPayoutDTO struct {
	FullName      string `json:"full_name"`
	PaypalEmail   string `json:"paypal_email"`
	DefaultMethod bool   `json:"default_method"`
	CreatorID     int
}
