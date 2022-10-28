package library

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SIBPayload struct {
	Sender      SIBSender `json:"sender"`
	To          []SIBTo   `json:"to"`
	Subject     string    `json:"subject"`
	HTMLContent string    `json:"htmlContent"`
}
type SIBSender struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type SIBTo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func SIBSendEmail(subject string, to SIBTo, sender SIBSender, html string) {
	data := SIBPayload{
		Sender:      sender,
		To:          []SIBTo{to},
		Subject:     subject,
		HTMLContent: html,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api.sendinblue.com/v3/smtp/email", body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Api-Key", os.Getenv("SIB_api_key"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
