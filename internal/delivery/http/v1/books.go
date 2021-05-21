package v1

import (
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
	fmt.Print("getBookById")
}

func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("update")
}

func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print("delete")
}
