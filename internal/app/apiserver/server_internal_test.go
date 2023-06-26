package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/pvelp/http-rest-api-template/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases := []struct {
		name         string
		payload      interface{}
		exceptedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"name":     "Pvelp",
				"password": "password",
			},
			exceptedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			exceptedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"name": "a",
			},
			exceptedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.exceptedCode, rec.Code)
		})
	}
	//rec := httptest.NewRecorder()
	//req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	//s := newServer(teststore.New())
	//s.ServeHTTP(rec, req)
	//assert.Equal(t, rec.Code, http.StatusOK)
}
