package sqlstore_test

import (
	"github.com/pvelp/http-rest-api-template/internal/app/model"
	"github.com/pvelp/http-rest-api-template/internal/app/store"
	"github.com/pvelp/http-rest-api-template/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dbUrl)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByCardId(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dbUrl)
	defer teardown("users")
	s := sqlstore.New(db)

	cardId := 0x1233
	_, err := s.User().FindByCardId(cardId)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.CardId = cardId
	s.User().Create(u)
	cardId = 0x123f
	u, err = s.User().FindByCardId(cardId)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
