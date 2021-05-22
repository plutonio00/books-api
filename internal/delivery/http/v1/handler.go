package v1

import (
	"github.com/gorilla/mux"
	"github.com/plutonio00/books-api/internal/service"
)

type Handler struct {
	baseEndpoint string
	services     *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		baseEndpoint: "/api/v1/",
		services:     services,
	}
}

func (h *Handler) Init(router *mux.Router) {
	h.initBooksRoutes(router)
	//     h.initAuthorsRoutes(router)
}
