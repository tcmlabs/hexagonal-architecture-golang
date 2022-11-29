package http

import (
	"encoding/json"
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"tcmlabs.fr/hexagonal_architecture_golang/internal/user/core/domain"
)

func getUser(userSvc services.User) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userSvc.Get()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		usersResponse := make([]User, len(users))
		for i, user := range users {
			userID, err := uuid.Parse(user.ID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			usersResponse[i] = User{
				Id:    userID,
				Email: openapi_types.Email(user.Email),
			}
		}

		err = json.NewEncoder(w).Encode(usersResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
