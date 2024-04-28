package endpoints

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/domain/entity"
)

func MakeCreateImageEndpoint(imageService *domain.ImageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body entity.Image

		err := r.ParseMultipartForm(200 << 20)
		if err != nil {
			http.Error(w, "Failed to parse multi form data", http.StatusInternalServerError)
		}

		id := r.FormValue("id")
		name := r.FormValue("name")
		dateString := r.FormValue("date")
		date, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("content")
		if err != nil {
			http.Error(w, "Failed to parse multipart form data", http.StatusBadRequest)
			return
		}
		defer file.Close()

		imageContent, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Failed to read image content", http.StatusInternalServerError)
			return
		}

		image := &entity.Image{
			Id:      id,
			Name:    name,
			Date:    date,
			Content: imageContent,
		}

		err = imageService.Register(image)
		if err != nil {
			http.Error(w, "Failed to create image", http.StatusInternalServerError)
			return
		}

		res := map[string]string{"image_id": body.Id}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to enconde response", http.StatusInternalServerError)
			return
		}
	}
}
