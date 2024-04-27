package rest

import (
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/transport/http/endpoints"
	"github.com/gorilla/mux"
)

type TestHandler struct {
	imageService *domain.ImageService
}

func NewTestHandler(imageService *domain.ImageService) *TestHandler {
	return &TestHandler{
		imageService: imageService,
	}
}

func (h *TestHandler) Register(router *mux.Router) {
	testEndpoint := endpoints.MakeTestEndpoint(h.imageService)

	router.HandleFunc("/test", testEndpoint).Methods(http.MethodGet)
}
