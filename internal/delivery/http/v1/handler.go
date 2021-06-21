package v1

import (
	"github.com/gorilla/mux"
	"github.com/plutonio00/books-api/internal/service"
	"github.com/plutonio00/books-api/pkg/token"
	"net/url"
)

type Handler struct {
	baseEndpoint string
	services     *service.Services
	tokenManager *token.TokenManager
}

func NewHandler(services *service.Services, tokenManager *token.TokenManager) *Handler {
	return &Handler{
		baseEndpoint: "/api/v1/",
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(router *mux.Router) {
	h.initBooksRoutes(router)
	h.initUsersRoutes(router)
}

func (h *Handler) ParsePostParamsToKeyValueSlices(params url.Values) ([]string, []interface{}) {
	keys := make([]string, 0)
	values := make([]interface{}, 0)

	for key, value := range params {
		keys = append(keys, key)
		values = append(values, value[0])
	}

	return keys, values
}
