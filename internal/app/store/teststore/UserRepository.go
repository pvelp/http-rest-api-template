package teststore

import (
	"github.com/pvelp/http-rest-api-template/internal/app/model"
	"github.com/pvelp/http-rest-api-template/internal/app/store"
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
		return nil, store.ErrRecordNotFound
	}
	return u, nil

}
