package endpoints

import (
	"fmt"
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
)

func MakeTestEndpoint(testService *domain.TestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := testService.Test()
		fmt.Printf("RESPONSE: %s", response)
		http.Error(w, "Endpoint not implemented yet", http.StatusInternalServerError)
	}
}
