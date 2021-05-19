package v1

import (
	"github.com/gorilla/mux"
)

type Handler struct {
	baseEndpoint string
}

func NewHandler() *Handler {
	return &Handler{
		baseEndpoint: "/api/v1/",
	}
}

func (h *Handler) Init(router *mux.Router) {
	h.initBooksRoutes(router)
	//     h.initAuthorsRoutes(router)
}
