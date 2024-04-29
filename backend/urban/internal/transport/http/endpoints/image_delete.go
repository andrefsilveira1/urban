package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/gorilla/mux"
)

func MakeDeleteImageEndpoint(imageService *domain.ImageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("Id received:", id)

		err := imageService.Delete(id)
		if err != nil {
			http.Error(w, "Failed to delete image", http.StatusInternalServerError)
			return
		}

		res := map[string]string{"image_id": id}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to enconde response", http.StatusInternalServerError)
			return
		}
	}
}
