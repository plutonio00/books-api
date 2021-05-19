package http

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	v1 "github.com/plutonio00/books-api/internal/delivery/http/v1"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() *mux.Router {
	router := mux.NewRouter()
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *mux.Router) {
	handlerV1 := v1.NewHandler()

	handlerV1.Init(router)
}
