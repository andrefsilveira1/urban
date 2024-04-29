package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
)

func MakeListImageEndpoint(imageService *domain.ImageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		image, err := imageService.List()
		if err != nil {
			fmt.Println("Error", err)
			http.Error(w, "Failed to list images", http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(image)
		if err != nil {
			http.Error(w, "Failed to serialize images", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
