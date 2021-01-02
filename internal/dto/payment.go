package dto

type PaymentDTO struct {
	Amount            float32 `json:"amount"`
	SenderID          int     `json:"sender_id"`
	ReceiverProjectID int     `json:"receiver_project_id"`
}

type DonationDTO struct {
	PaymentDTO
	Message string `json:"message"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
}
