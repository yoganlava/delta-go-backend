package entity

import "time"

type Subscription struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	ExpiryDate time.Time `json:"expiry_date"`
	TierID     int       `json:"tier_id"`
	CreatedAt  time.Time `json:"created_at"`
	WillRenew  bool      `json:"will_renew"`
	Meta       SubscriptionMeta
}

type SubscriptionMeta struct {
	StripeSubscriptionID *string `json:"stripe_subscription_id,omitempty"`
	PaypalSubscriptionID *string `json:"paypal_subscription_id,omitempty"`
}
