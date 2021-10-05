package main

import (
	"errors"
	"fmt"
)

type User struct {
	email		string
	username	string
	password	string
}

type Authenticator struct {

}

func (a Authenticator) Authenticate(adapter AuthAdapter, user User) error {
	result := adapter.Authenticate(user)
	if result {
		return nil
	}

	return errors.New("invalid credentials")
}

type AuthAdapter interface {
	Authenticate(user User) bool
}

type GoogleAuth struct {

}

func (g GoogleAuth) Authenticate(username string, password string) string {
	if username == "admin" && password == "123456" {
		return "SUCCESS"
	}

	return "FAIL"
}

type GoogleAuthAdapter struct {
	auth GoogleAuth
}

func (g GoogleAuthAdapter) Authenticate(user User) bool {
	return g.auth.Authenticate(user.username, user.password) == "SUCCESS"
}

type MicrosoftAuth struct {

}

func (m MicrosoftAuth) Authenticate(email string, password string) bool {
	if email == "admin@microsoft.com" && password == "123456" {
		return true
	}

	return false
}

type MicrosoftAuthAdapter struct {
	auth MicrosoftAuth
}

func (m MicrosoftAuthAdapter) Authenticate(user User) bool {
	return m.auth.Authenticate(user.email, user.password)
}

func main() {
	authenticator := Authenticator{}

	googleAuth := GoogleAuth{}
	googleAuthAdapter := GoogleAuthAdapter{googleAuth}
	googleUser := User{"admin@google.com", "admin", "123456"}

	microsoftAuth := MicrosoftAuth{}
	microsoftAuthAdapter := MicrosoftAuthAdapter{microsoftAuth}
	microsoftUser := User{"admin@microsoft.com", "admin", "123456"}

	if err := authenticator.Authenticate(googleAuthAdapter, googleUser); err != nil {
		fmt.Println("google auth fail:", err)
	} else {
		fmt.Println("google auth success")
	}

	if err := authenticator.Authenticate(microsoftAuthAdapter, microsoftUser); err != nil {
		fmt.Println("microsoft auth fail:", err)
	} else {
		fmt.Println("microsoft auth success")
	}
}