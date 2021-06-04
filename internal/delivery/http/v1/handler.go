package v1

import (
	"github.com/gorilla/mux"
	"github.com/plutonio00/books-api/internal/service"
	"net/url"
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

func (h *Handler) ParsePostParamsToKeyValueSlices(params url.Values) ([]string, []interface{}) {
    keys := make([]string, 0)
    values := make([]interface{}, 0)

    for key, value := range params {
    	keys = append(keys, key)
    	values = append(values, value[0])
    }

    return keys, values
}
