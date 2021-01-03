package paypal

import (
	"context"
	"main/db"
	"main/internal/dto"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/plutov/paypal"
)

type IPaypalService interface {
	CreatePaypalPayout(createPaypalPayoutDTO dto.CreatePaypalPayoutDTO) error
}

type PaypalService struct {
	pool   *pgxpool.Pool
	client *paypal.Client
}

func New() PaypalService {
	c, err := paypal.NewClient(os.Getenv("PAYPAL_CLIENT"), os.Getenv("PAYPAL_SECRET"), paypal.APIBaseSandBox)
	if err != nil {
		panic(err)
	}
	return PaypalService{db.Connection(), c}
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

func (ps PaypalService) CreatePaypalOrder(createDonationPaypalDTO dto.CreateDonationPaypalOrderDTO) (*paypal.Order, error) {
	return ps.client.CreateOrder("CAPTURE", []paypal.PurchaseUnitRequest{
		paypal.PurchaseUnitRequest{
			Amount: &paypal.PurchaseUnitAmount{
				Value:    createDonationPaypalDTO.Amount,
				Currency: createDonationPaypalDTO.Currency,
			},
			Payee: &paypal.PayeeForOrders{
				EmailAddress: createDonationPaypalDTO.CreatorEmail,
			},
		},
	},
		&paypal.CreateOrderPayer{
			Name: &paypal.CreateOrderPayerName{
				GivenName: createDonationPaypalDTO.PayerFirstName,
				Surname:   createDonationPaypalDTO.PayerLastName,
			},
			EmailAddress: createDonationPaypalDTO.PayerEmail,
		},
		&paypal.ApplicationContext{
			BrandName:          "",
			Locale:             "",
			LandingPage:        "",
			ShippingPreference: "",
			UserAction:         "",
			ReturnURL:          "http://onjin.jp",
			CancelURL:          "http://onjin.jp",
		},
	)
}
