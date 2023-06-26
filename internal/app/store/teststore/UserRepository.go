package teststore

import (
	"errors"
	"github.com/pvelp/http-rest-api-template/internal/app/model"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.CardId] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByCardId(cardId int) (*model.User, error) {
	u, ok := r.users[cardId]
	if !ok {
		return nil, errors.New("Not found")
	}
	return u, nil

}
