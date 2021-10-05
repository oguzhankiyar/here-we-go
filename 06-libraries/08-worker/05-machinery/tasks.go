package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

type SendMailPayload struct {
	Email   string `json:"email"`
}

func SendMail(msg string) (bool, error) {
	payload := SendMailPayload{}

	bytes, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(bytes, &payload)
	if err != nil {
		return false, err
	}

	if len(payload.Email) == 0 {
		return false, errors.New("invalid task")
	}

	fmt.Println("sending mail to ", payload.Email)

	return true, nil
}
