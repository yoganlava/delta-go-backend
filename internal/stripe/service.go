package stripe

import (
	"github.com/stripe/stripe-go/account"
	"github.com/stripe/stripe-go/v72"
)

func handleSubscriptionWebhook() {

}

func CreateStripeAccount(email string) (*stripe.Account, error) {
	params := &stripe.AccountParams{
		Country: stripe.String("JP"),
		Email:   stripe.String(email),
		Type:    stripe.String(string(stripe.AccountTypeStandard)),
	}
	acct, err := account.New(params)
	return acct, err
}

func handleDonation() {

}
