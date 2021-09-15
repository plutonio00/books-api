package v1

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	api_errors "github.com/plutonio00/books-api/internal/error"
	"github.com/plutonio00/books-api/internal/model"
	"net/http"
	"strconv"
)

func (h *Handler) initBooksRoutes(router *mux.Router) {
	booksRouter := router.PathPrefix(h.baseEndpoint + "books").Subrouter()
	booksRouter.Use(h.authWithToken)
	booksRouter.HandleFunc("/list", h.getBooksList).Methods("GET")
	booksRouter.HandleFunc("/{id:[0-9]+}", h.getBookById).Methods("GET")
	booksRouter.HandleFunc("/update", h.updateBook).Methods("POST")
	booksRouter.HandleFunc("/delete/{id:[0-9]+}", h.deleteBook).Methods("DELETE")
}

// @Summary Books list
// @Tags books
// @Description endpoint for getting all books
// @Produce json
// @Security ApiTokenAuth
// @Success 200 {object} ApiResponse{result=[]model.Book}
// @Failure 400 {object} ApiResponse{result=string}
// @Failure 500 {object} ApiResponse{result=string}
// @Failure default {object} ApiResponse{result=string}
// @Router /books/list [get]
func (h *Handler) getBooksList(w http.ResponseWriter, r *http.Request) {
	data, err := h.services.Books.GetBooksList()

	if err != nil {
		if err == api_errors.ErrBookNotFound {
			jsonResponse(w, http.StatusNotFound, err.Error())
			return
		}

		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, data)
}

// @Summary Books by id
// @Tags books
// @Description endpoint for getting book by id
// @Produce json
// @Security ApiTokenAuth
// @Param id path integer true "Book's id"
// @Success 200 {object} ApiResponse{result=model.Book}
// @Failure 400 {object} ApiResponse{result=string}
// @Failure 404 {object} ApiResponse{result=string}
// @Failure 500 {object} ApiResponse{result=string}
// @Failure default {object} ApiResponse{result=string}
// @Router /books/{id} [get]
func (h *Handler) getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book, err := h.services.Books.FindById(id)

	if err != nil {
		if book == nil {
			jsonResponse(w, http.StatusNotFound, err.Error())
			return
		}

		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, book)
	return
}

// @Summary Update book by id
// @Tags books
// @Description endpoint for update book by id
// @Produce json
// @Security ApiTokenAuth
// @Param id formData integer true "Book's id"
// @Param title formData string false "Book's title"
// @Param releaseYear formData integer false "Book's release year"
// @Param authorId formData integer false "Book's author"
// @Success 200 {object} ApiResponse{result=string}
// @Failure 400 {object} ApiResponse{result=string}
// @Failure 404 {object} ApiResponse{result=string}
// @Failure 500 {object} ApiResponse{result=string}
// @Failure default {object} ApiResponse{result=string}
// @Router /books/update [post]
func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var book model.Book
	var decoder = schema.NewDecoder()

	err := decoder.Decode(&book, r.PostForm)

	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	//err = book.Validate()

	//     if err != nil {
	//         jsonResponse(w, http.StatusBadRequest, err.Error())
	//         return
	//     }

	err = h.services.Books.Update(&book)

	if err != nil {
		if err == api_errors.ErrBookNotFound {
			jsonResponse(w, http.StatusNotFound, err.Error())
			return
		}

		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, "Books were updated successfully")
	return
}

// @Summary Delete book by id
// @Tags books
// @Description endpoint for delete book by id
// @Produce json
// @Security ApiTokenAuth
// @Param id path integer true "Book's id"
// @Success 200 {object} ApiResponse{result=string}
// @Failure 400 {object} ApiResponse{result=string}
// @Failure 404 {object} ApiResponse{result=string}
// @Failure 500 {object} ApiResponse{result=string}
// @Failure default {object} ApiResponse{result=string}
// @Router /books/delete/{id} [delete]
func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := h.services.Books.DeleteById(id)

	if err != nil {
		if err == api_errors.ErrBookNotFound {
			jsonResponse(w, http.StatusNotFound, err.Error())
			return
		}

		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, fmt.Sprintf("Book with id %d was deleted", id))
	return
}
