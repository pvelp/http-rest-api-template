package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Name              string
	Surname           string
	Password          string
	EncryptedPassword string
	CardId            int
	IsWorker          bool
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u, validation.Field(&u.Name, validation.Required, validation.Length(2, 32)))
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
