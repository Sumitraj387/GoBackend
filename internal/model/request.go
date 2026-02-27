package model

import (
	"errors"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Phone string `json:"phone_number"`
	Email string `json:"email_id"`
}

func (user *User) Validate() error {
	if len(strings.TrimSpace(user.Name)) == 0 {
		return errors.New("user name must not be empty")
	}
	if len(strings.TrimSpace(user.Phone)) == 0 {
		return errors.New("phone no must not be empty")
	}
	if len(strings.TrimSpace(user.Email)) == 0 {
		return errors.New("user email must not be empty")
	}
	return nil
}
