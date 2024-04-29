package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gorilla/mux"
)

func MakeGetUserEndpoint(userService *domain.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("Id received:", id)

		user, err := userService.Get(id)
		if err != nil {
			errorResponse := entity.ErrorResponse{Message: fmt.Sprintf("Failed to delete user: %v", err)}
			responseJSON, _ := json.Marshal(errorResponse)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJSON)
			return
		}
		response, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Failed to serialize user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
