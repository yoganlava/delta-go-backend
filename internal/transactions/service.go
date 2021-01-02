package transactions

import (
	"context"
	"main/db"
	"main/internal/dto"

	"github.com/jackc/pgx/v4/pgxpool"
)

type TransactionService struct {
	pool *pgxpool.Pool
}

type ITransactionService interface{}

func New() TransactionService {
	return TransactionService{db.Connection()}
}

func (ts TransactionService) AddDonationTransaction(donationTransactionDTO dto.DonationTransactionDTO) error {
	var id int
	err := ts.pool.QueryRow(context.Background(), "insert into transaction (amount, sender_id, receiver_project_id, created_at) values ($1, $2, $3, now())", donationTransactionDTO.Amount, donationTransactionDTO.SenderID, donationTransactionDTO.ReceiverProjectID).Scan(&id)
	if err != nil {
		return err
	}
	_, err = ts.pool.Exec(context.Background(), "insert into donation_transaction (transaction_id, message, name, private) values ($1, $2, $3, $4)", id, donationTransactionDTO.Message, donationTransactionDTO.Name, donationTransactionDTO.Private)
	return err
}
