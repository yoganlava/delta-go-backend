package dto

type CreatePaypalPayoutDTO struct {
	FullName      string `json:"full_name"`
	PaypalEmail   string `json:"paypal_email"`
	DefaultMethod bool   `json:"default_method"`
	CreatorID     int
}

type CreateDonationPaypalOrderDTO struct {
	PayerFirstName string
	PayerLastName  string
	PayerEmail     string
	Amount         string
	Currency       string
	CreatorEmail   string
}
