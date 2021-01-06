package dto

type TransactionDTO struct {
	Amount            float32 `json:"amount"`
	SenderID          int     `json:"sender_id,omitempty"`
	ReceiverProjectID int     `json:"receiver_project_id"`
}

type DonationTransactionDTO struct {
	TransactionDTO
	Message string `json:"message"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

type PayoutTransaction struct {
	TransactionDTO
	PayoutMethodID int
}

type SubscriptionTransaction struct {
	TransactionDTO
	SubscriptionID int
}
