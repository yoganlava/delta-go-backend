package entity

type SendInBlueSender struct {
	Name  string  `json:"name"`
	Email string  `json:"email"`
	ID    float32 `json:"id"`
}
type SendInBlueTo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type SendInBlueType struct {
	HTMLContent string           `json:"htmlContent"`
	Sender      SendInBlueSender `json:"sender"`
	To          []SendInBlueTo   `json:"to"`
}
