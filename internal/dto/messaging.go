package dto

type CreateMessageDTO struct {
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Subject    string `json:"subject"`
	Body       string `json:"body"`
}
