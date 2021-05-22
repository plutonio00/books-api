package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) initBooksRoutes(router *mux.Router) {
	groupEndpoint := h.baseEndpoint + "books/"

	router.HandleFunc(groupEndpoint+"list", h.getBooksList).Methods("GET")
	router.HandleFunc(groupEndpoint+"{id:[0-9]+}", h.getBookById).Methods("GET")
	router.HandleFunc(groupEndpoint+"update", h.updateBook).Methods("POST")
	router.HandleFunc(groupEndpoint+"delete", h.deleteBook).Methods("DELETE")
}

func (h *Handler) getBooksList(w http.ResponseWriter, r *http.Request) {
	fmt.Print("bookList")
}

func (h *Handler) getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := h.services.Books.FindById(id)

	if err != nil {
		message, _ := json.Marshal("Error: empty id")
		jsonResponse(w, http.StatusBadRequest, message)
		return
	}

	res, err := json.Marshal(data)

	if err != nil {
	    fmt.Println(err)
    	message, _ := json.Marshal("Error: empty id")
    	jsonResponse(w, http.StatusBadRequest, message)
    	return
    }

	jsonResponse(w, http.StatusOK, res)
}

func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("update")
}

func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("delete")
}
