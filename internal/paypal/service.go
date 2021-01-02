package paypal

import (
	"context"
	"main/internal/dto"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IPaypalService interface {
	CreatePaypalPayout(createPaypalPayoutDTO dto.CreatePaypalPayoutDTO) error
}

type PaypalService struct {
	pool *pgxpool.Pool
}

func (ps PaypalService) CreatePaypalPayout(createPaypalPayoutDTO dto.CreatePaypalPayoutDTO) error {
	var payout_method_id int
	err := ps.pool.QueryRow(context.Background(), `insert into payout_method (full_name, created_at, updated_at, creator_id, default_method) values ($1, now(), now(), $2, $3) returning id`, createPaypalPayoutDTO.FullName, createPaypalPayoutDTO.CreatorID, createPaypalPayoutDTO.DefaultMethod).Scan(&payout_method_id)
	if err != nil {
		return nil
	}
	_, err = ps.pool.Exec(context.Background(), `insert into paypal_payout_method (payout_method_id, paypal_email) values ($1, $2)`, payout_method_id, createPaypalPayoutDTO.PaypalEmail)

	return nil
}
