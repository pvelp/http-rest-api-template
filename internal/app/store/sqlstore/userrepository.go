package sqlstore

import (
	"database/sql"
	"github.com/pvelp/http-rest-api-template/internal/app/model"
	"github.com/pvelp/http-rest-api-template/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	if err := u.Validate(); err != nil {
		return err
	}
	return r.store.db.QueryRow(
		"INSERT INTO users (name, surname, email, encrypted_password, card_id, is_worker) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		u.Name, u.Surname, u.Email, u.EncryptedPassword, u.CardId, u.IsWorker,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByCardId(cardId int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, name, surname, email, encrypted_password, card_id, is_worker FROM users WHERE card_id=$1",
		cardId,
	).Scan(&u.ID,
		&u.Name,
		&u.Surname,
		&u.Email,
		&u.EncryptedPassword,
		&u.CardId,
		&u.IsWorker,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, name, surname, email, encrypted_password, card_id, is_worker FROM users WHERE email=$1",
		email,
	).Scan(&u.ID,
		&u.Name,
		&u.Surname,
		&u.Email,
		&u.EncryptedPassword,
		&u.CardId,
		&u.IsWorker,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
