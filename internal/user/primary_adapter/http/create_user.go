package http

import (
	"encoding/json"
	"log"
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"gitlab.com/tclaudel_ateme/hexagonal_architecture_golang/internal/user/core/services"
)

func createUser(userSvc services.User) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user CreateUser

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		dUser, err := userSvc.Create(string(user.Email))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userID, err := uuid.Parse(dUser.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseBody, err := json.Marshal(User{
			Id:    userID,
			Email: openapi_types.Email(dUser.Email),
		})

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write(responseBody)
		if err != nil {
			log.Print(err)
		}
		return
	}
}
