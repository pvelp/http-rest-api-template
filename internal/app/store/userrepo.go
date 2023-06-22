package store

import "github.com/pvelp/http-rest-api-template/internal/app/model"

type UserRepo struct {
	store *Store
}

func (r *UserRepo) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (name, surname, card_id, is_worker) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Name, u.Surname, u.CardId, u.IsWorker,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepo) FindByCardId(cardId int) (*model.User, error) {
	return nil, nil
}
