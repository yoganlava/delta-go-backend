package entity

type Transaction struct {
	Amount            float32 `json:"amount"`
	SenderID          int     `json:"sender_id,omitempty"`
	ReceiverProjectID int     `json:"receiver_project_id"`
}

type DonationTransaction struct {
	Transaction
	Message string `json:"message"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

type PayoutTransaction struct {
	Transaction
	PayoutMethodID int
}

type SubscriptionTransaction struct {
	Transaction
	SubscriptionID int
}
