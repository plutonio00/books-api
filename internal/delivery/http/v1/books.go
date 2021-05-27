package v1

import (
	"fmt"
	"github.com/gorilla/mux"
	api_errors "github.com/plutonio00/books-api/internal/error"
	"net/http"
)

func (h *Handler) initBooksRoutes(router *mux.Router) {
	groupEndpoint := h.baseEndpoint + "books/"

	router.HandleFunc(groupEndpoint+"list", h.getBooksList).Methods("GET")
	router.HandleFunc(groupEndpoint+"{id:[0-9]+}", h.getBookById).Methods("GET")
	router.HandleFunc(groupEndpoint+"update", h.updateBook).Methods("POST")
	router.HandleFunc(groupEndpoint+"delete/{id:[0-9]+}", h.deleteBook).Methods("DELETE")
}

func (h *Handler) getBooksList(w http.ResponseWriter, r *http.Request) {
	data, err := h.services.Books.GetBooksList()

	if err != nil {
		if err == api_errors.ErrBookNotFound {
			jsonResponse(w, http.StatusNotFound, err.Error())
		}

		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, data)
}

func (h *Handler) getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := h.services.Books.FindById(id)

	if err != nil {
		if err == api_errors.ErrBookNotFound {
			jsonResponse(w, http.StatusNotFound, err.Error())
		}

		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, data)
	return
}

func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	params := r.Form

	id := r.Form.Get("id")

	if id == "" {
		jsonResponse(w, http.StatusBadRequest, "Missing mandatory parameter 'id'")
		return
	}

	keys := make([]string, 0)
	values := make([]interface{}, 0)

	for key, value := range params {
		keys = append(keys, key)
		values = append(values, value[0])
	}

	values = append(values, id)
	err := h.services.Books.UpdateById(keys, values)

	if err != nil {
		if err == api_errors.ErrBookNotFound {
			jsonResponse(w, http.StatusNotFound, err.Error())
		}

		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("Book with id %s updated", id))
	return
}

func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.services.Books.DeleteById(id)

	if err != nil {
		if err == api_errors.ErrBookNotFound {
			jsonResponse(w, http.StatusNotFound, err.Error())
		}

		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("Book with id %s deleted", id))
	return
}
