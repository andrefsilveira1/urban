package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/domain/entity"
)

func MakeListImageEndpoint(imageService *domain.ImageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		image, err := imageService.List()
		if err != nil {
			errorResponse := entity.ErrorResponse{Message: fmt.Sprintf("Failed to delete user: %v", err)}
			responseJSON, _ := json.Marshal(errorResponse)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(responseJSON)
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
