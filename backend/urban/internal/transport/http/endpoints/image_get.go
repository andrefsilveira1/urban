package endpoints

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/gorilla/mux"
)

func MakeGetImageEndpoint(imageService *domain.ImageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("ID:", id)

		image, err := imageService.Get(id)

		if err != nil {
			http.Error(w, "Failed to decode image content", http.StatusInternalServerError)
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(string(image.Content))
		if err != nil {
			http.Error(w, "Failed to decode image string", http.StatusInternalServerError)
			return
		}

		filepath := "decoded_image.png"
		err = os.WriteFile(filepath, decoded, 0644)
		if err != nil {
			fmt.Println("Error writing decoded content to file:", err)
			return
		}
		fmt.Printf("Decoded image saved as: %s\n", filepath)

		if err != nil {
			http.Error(w, "Failed to retrieve image", http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(image)
		if err != nil {
			http.Error(w, "Failed to serialize image", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
