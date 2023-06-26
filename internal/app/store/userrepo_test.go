package store_test

import (
	"github.com/pvelp/http-rest-api-template/internal/app/model"
	"github.com/pvelp/http-rest-api-template/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	s, teardown := store.TestStore(t, dbUrl)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByCardId(t *testing.T) {
	s, teardown := store.TestStore(t, dbUrl)
	defer teardown("users")

	cardId := 0x123f
	_, err := s.User().FindByCardId(cardId)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.CardId = cardId
	s.User().Create(u)

	u, err = s.User().FindByCardId(cardId)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
