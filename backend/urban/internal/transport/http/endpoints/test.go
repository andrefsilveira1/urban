package endpoints

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/andrefsilveira1/urban/internal/domain"
)

func MakeTestEndpoint(testService *domain.TestService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := testService.Test()
		fmt.Printf("RESPONSE: %s", response)
		w.WriteHeader(http.StatusOK)
		res := map[string]string{"status": "connected"}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to enconde response", http.StatusInternalServerError)
			return
		}
	}
}
