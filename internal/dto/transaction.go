package dto

type TransactionDTO struct {
	Amount            float32
	SenderID          int
	ReceiverProjectID int
}

type DonationTransactionDTO struct {
	TransactionDTO
	Message string
	Name    string
	Private bool
}
