package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gorilla/mux"
)

func MakeDeleteUserEndpoint(userService *domain.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("Id received:", id)

		err := userService.Delete(id)
		if err != nil {
			errorResponse := entity.ErrorResponse{Message: fmt.Sprintf("Failed to delete user: %v", err)}
			responseJSON, _ := json.Marshal(errorResponse)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJSON)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		res := map[string]string{"image_id": id}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to enconde response", http.StatusInternalServerError)
			return
		}
	}
}
