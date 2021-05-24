package v1

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/plutonio00/books-api/internal/service"
)

func (h *Handler) initBooksRoutes(router *mux.Router) {
	groupEndpoint := h.baseEndpoint + "books/"

	router.HandleFunc(groupEndpoint+"list", h.getBooksList).Methods("GET")
	router.HandleFunc(groupEndpoint+"{id:[0-9]+}", h.getBookById).Methods("GET")
	router.HandleFunc(groupEndpoint+"update", h.updateBook).Methods("POST")
	router.HandleFunc(groupEndpoint+"delete", h.deleteBook).Methods("DELETE")
}

func (h *Handler) getBooksList(w http.ResponseWriter, r *http.Request) {
	data, err := h.services.Books.GetBooksList()

	if err != nil {
    	    if err == service.ErrBookNotFound {
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
	    if err == service.ErrBookNotFound {
	        jsonResponse(w, http.StatusNotFound, err.Error())
	    }

		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, data)
}

func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("update")
}

func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("delete")
}
