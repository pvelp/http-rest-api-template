package teststore_test

import (
	"github.com/pvelp/http-rest-api-template/internal/app/model"
	"github.com/pvelp/http-rest-api-template/internal/app/store"
	"github.com/pvelp/http-rest-api-template/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByCardId(t *testing.T) {
	s := teststore.New()
	cardId := 0x1233
	_, err := s.User().FindByCardId(cardId)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	cardId = 0x123f
	u := model.TestUser(t)
	u.CardId = cardId
	s.User().Create(u)

	u, err = s.User().FindByCardId(cardId)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
