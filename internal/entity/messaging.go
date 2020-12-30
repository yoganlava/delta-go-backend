package entity

type Message struct {
	id         int
	SenderID   int
	ReceiverID int
	Subject    string
	Body       string
	CreatedAt  string
}
