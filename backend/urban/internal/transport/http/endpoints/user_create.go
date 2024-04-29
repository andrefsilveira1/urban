package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gocql/gocql"
)

func MakeCreateUserEndpoint(userService *domain.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		user := &entity.User{
			Id:       gocql.TimeUUID(),
			Name:     name,
			Email:    email,
			Password: password,
		}

		err := userService.Register(user)
		if err != nil {
			errorResponse := entity.ErrorResponse{Message: fmt.Sprintf("Failed to delete user: %v", err)}
			responseJSON, _ := json.Marshal(errorResponse)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJSON)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		res := map[string]gocql.UUID{"user_uid": user.Id}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to create image", http.StatusInternalServerError)
			return
		}
	}
}
