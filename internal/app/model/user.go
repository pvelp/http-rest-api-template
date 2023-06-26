package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                int
	Name              string
	Surname           string
	Password          string
	EncryptedPassword string
	CardId            int
	IsWorker          bool
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
