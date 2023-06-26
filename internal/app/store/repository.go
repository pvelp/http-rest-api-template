package store

import "github.com/pvelp/http-rest-api-template/internal/app/model"

type UserRepository interface {
	Create(user *model.User) error
	//FindByCardId(cardId int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
