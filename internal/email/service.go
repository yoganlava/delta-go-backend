package email

import (
	"bytes"
	"encoding/json"
	"main/internal/entity"
	"net/http"
	"os"
)

func SendEmail(email string, username string, subject string, body string) error {
	postBody, _ := json.Marshal(entity.SendInBlueType{
		Sender: entity.SendInBlueSender{
			Name:  "Onjin",
			Email: "noreply@onjin.jp",
			ID:    -2,
		},
		To: []entity.SendInBlueTo{{
			Email: email,
			Name:  username,
		}},
		Subject:     subject,
		HTMLContent: body,
	})
	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", "https://api.sendinblue.com/v3/smtp/email", responseBody)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", os.Getenv("EMAIL_API_KEY"))

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
