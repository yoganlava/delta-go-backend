package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/internal/entity"
	"net/http"
	"os"
)

func SendVerificationEmail(email string, username string, token string) error {
	emailHTML := fmt.Sprintf(
		`<h1>Welcome %v</h1>
	<a href='http://onjin.jp/verify/%v'>Click me</a>
	`, username, token)
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
		HTMLContent: emailHTML,
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
