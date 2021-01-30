package entity

import "time"

type PayoutMonth struct {
	Amount float32   `json:"amount"`
	Date   time.Time `json:"date"`
}
