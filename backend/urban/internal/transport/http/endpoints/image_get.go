package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gorilla/mux"
)

func MakeGetImageEndpoint(imageService *domain.ImageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Println("Id received:", id)

		image, err := imageService.Get(id)
		// testImage(image.Content)

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
			http.Error(w, "Failed to serialize image", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// func testImage(content []byte) {
// 		encode := base64.StdEncoding.EncodeToString(content)
// 		decoded, err := base64.StdEncoding.DecodeString(encode)
// 		if err != nil {
// 			fmt.Println("Error to decode image")
// 			return
// 		}

// 		filepath := "decoded_image.png"
// 		err = os.WriteFile(filepath, decoded, 0644)
// 		if err != nil {
// 			fmt.Println("Error writing decoded content to file:", err)
// 			return
// 		}
// 		fmt.Printf("Decoded image saved as: %s\n", filepath)

// 		if err != nil {
// 			fmt.Println("Error to read image")
// 			return
// 		}
// }
