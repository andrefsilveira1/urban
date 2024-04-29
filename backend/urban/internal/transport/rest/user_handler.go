package rest

import (
	"net/http"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/transport/http/endpoints"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService *domain.UserService
}

func NewUserHandler(userService *domain.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(router *mux.Router) {
	createUserEndpoint := endpoints.MakeCreateUserEndpoint(h.userService)
	getUserEndpoint := endpoints.MakeGetUserEndpoint(h.userService)
	deleteUserEndpoint := endpoints.MakeDeleteUserEndpoint(h.userService)

	router.HandleFunc("/users", createUserEndpoint).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", getUserEndpoint).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", deleteUserEndpoint).Methods(http.MethodDelete)
}
