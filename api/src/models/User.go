package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) FormatAndValidate(operation string) error {
	user.format(operation)
	if error := user.validate(operation); error != nil {
		return error
	}
	return nil
}

func (user *User) format(operation string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if operation == "create" {
		passwordHash, error := security.GeneratedHash(user.Password)
		if error != nil {
			return error
		}
		user.Password = string(passwordHash)
	}
	return nil
}

func (user *User) validate(operation string) error {
	if user.Name == "" {
		return errors.New("The name is mandatory e cannot be empty or blank")
	}
	if user.Nick == "" {
		return errors.New("The nick is mandatory e cannot be empty or blank")
	}
	if user.Email == "" {
		return errors.New("The email is mandatory e cannot be empty or blank")
	}
	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return error
	}
	if operation == "create" && user.Password == "" {
		return errors.New("The password is mandatory e cannot be empty or blank")
	}
	return nil
}
