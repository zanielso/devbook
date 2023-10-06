package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) FormatAndValidate() error {
	user.format()
	if error := user.validate(); error != nil {
		return error
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("The name is mandatory e cannot be empty or blank")
	}
	if user.Nick == "" {
		return errors.New("The nick is mandatory e cannot be empty or blank")
	}
	if user.Email == "" {
		return errors.New("The email is mandatory e cannot be empty or blank")
	}
	if user.Password == "" {
		return errors.New("The password is mandatory e cannot be empty or blank")
	}
	return nil
}
