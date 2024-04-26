package endpoints

import (
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
)

func MakeCreateUserEndpoint(userService *domain.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Endpoint not implemented yet", http.StatusInternalServerError)
	}
}