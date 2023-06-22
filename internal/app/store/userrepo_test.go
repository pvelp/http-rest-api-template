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

	u, err := s.User().Create(&model.User{
		Name:     "Nikolay",
		Surname:  "Krapivyansky",
		CardId:   0x134a,
		IsWorker: false,
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
