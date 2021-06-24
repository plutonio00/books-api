package http

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/plutonio00/books-api/docs"
	"github.com/plutonio00/books-api/internal/config"
	v1 "github.com/plutonio00/books-api/internal/delivery/http/v1"
	"github.com/plutonio00/books-api/internal/service"
	"github.com/plutonio00/books-api/pkg/token"
	"github.com/swaggo/http-swagger"
)

type Handler struct {
	services     *service.Services
	tokenManager *token.TokenManager
}

func NewHandler(services *service.Services, tokenManager *token.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(conf *config.Config) *mux.Router {
	router := mux.NewRouter()
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *mux.Router) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	handlerV1.Init(router)
}
