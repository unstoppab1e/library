package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unstoppab1e/library/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "name",
			Email:     "uses",
			Password:  "password",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusBadRequest)
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
