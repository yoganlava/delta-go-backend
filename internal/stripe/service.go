package stripe

import (
	"context"
	"main/internal/dto"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stripe/stripe-go/account"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/v72"
)

func handleSubscriptionWebhook() {

}

type StripeService struct {
	pool *pgxpool.Pool
}

type IStripeService interface {
}

func New()

func CreateStripeAccount(email string) (*stripe.Account, error) {
	params := &stripe.AccountParams{
		Country: stripe.String("JP"),
		Email:   stripe.String(email),
		Type:    stripe.String(string(stripe.AccountTypeStandard)),
	}
	acct, err := account.New(params)
	return acct, err
}

func (ss StripeService) CreatePaymentIntent(donationDTO dto.DonationDTO) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Amount:               stripe.Int64(int64(donationDTO.Amount)),
		Currency:             stripe.String(string(stripe.CurrencyJPY)),
		ApplicationFeeAmount: stripe.Int64(0),
	}
	var account_id string
	err := ss.pool.QueryRow(context.Background(), `select creator.stripe_account_id from creator inner join project on creator.id=(select creator_id from project where id=$1)`, donationTransactionDTO).Scan(&account_id)
	if err != nil {
		return nil, err
	}
	params.SetStripeAccount(account_id)
	pi, err := paymentintent.New(params)

	return pi, err
}
