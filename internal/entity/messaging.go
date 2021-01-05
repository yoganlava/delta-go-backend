package entity

import "time"

type Message struct {
	id         int
	SenderID   int
	ReceiverID int
	Subject    string
	Body       string
	CreatedAt  time.Time
}
