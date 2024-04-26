package rest

import (
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/transport/http/endpoints"
	"github.com/gorilla/mux"
)

type ImageHandler struct {
	imageService *domain.ImageService
}

func NewImageHandler(imageService *domain.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}

func (h *ImageHandler) Register(router *mux.Router) {
	createImageEndpoint := endpoints.MakeCreateImageEndpoint(h.imageService)
	getImageEndpoint := endpoints.MakeGetImageEndpoint(h.imageService)
	listImageEndpoint := endpoints.MakeListImageEndpoint(h.imageService)
	deleteImageEndpoint := endpoints.MakeDeleteImageEndpoint(h.imageService)

	router.HandleFunc("/images", createImageEndpoint).Methods(http.MethodPost)
	router.HandleFunc("/images/{id}", getImageEndpoint).Methods(http.MethodGet)
	router.HandleFunc("/images", listImageEndpoint).Methods(http.MethodGet)
	router.HandleFunc("/images/{id}", deleteImageEndpoint).Methods(http.MethodDelete)
	// protected := router.Pathprefix("/").Subrouter()
}
