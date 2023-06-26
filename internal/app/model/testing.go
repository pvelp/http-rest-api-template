package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Name:     "Test",
		Surname:  "Surtest",
		Email:    "example@example.com",
		Password: "password",
		CardId:   0x1234,
		IsWorker: false,
	}
}
