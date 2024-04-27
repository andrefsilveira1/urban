package rest

import (
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/transport/http/endpoints"
	"github.com/gorilla/mux"
)

type TestHandler struct {
	testService *domain.TestService
}

func NewTestHandler(testService *domain.TestService) *TestHandler {
	return &TestHandler{
		testService: testService,
	}
}

func (h *TestHandler) Register(router *mux.Router) {
	testEndpoint := endpoints.MakeTestEndpoint(h.testService)

	router.HandleFunc("/test", testEndpoint).Methods(http.MethodGet)
}
