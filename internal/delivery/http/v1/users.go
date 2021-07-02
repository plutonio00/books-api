package v1

import (
	"github.com/gorilla/mux"
// 	api_errors "github.com/plutonio00/books-api/internal/error"
	"net/http"
	"github.com/gorilla/schema"
	"github.com/plutonio00/books-api/internal/model/input"
)

func (h *Handler) initUsersRoutes(router *mux.Router) {
	usersRouter := router.PathPrefix(h.baseEndpoint + "users").Subrouter()

	usersRouter.HandleFunc("/sign-up", h.SignUp).Methods("POST")
	usersRouter.HandleFunc("/sign-in", h.SignIn).Methods("POST")
}

// @Summary User sign in
// @Tags users
// @Description endpoint for user sign in
// @Produce json
// @Param email formData string true "User's email"
// @Param password formData string true "User's password"
// @Success 200 {object} ApiResponse{result=service.Token}
// @Failure 400 {object} ApiResponse{result=string}
// @Failure 404 {object} ApiResponse{result=string}
// @Failure 500 {object} ApiResponse{result=string}
// @Failure default {object} ApiResponse{result=string}
// @Router /users/sign-in [post]
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var credentials input.UserCredentials
	var decoder = schema.NewDecoder()

	err := decoder.Decode(&credentials, r.PostForm)

	if err != nil {
	    jsonResponse(w, http.StatusInternalServerError, err.Error())
    	return
	}

// 	if email == "" || password == "" {
// 		jsonResponse(w, http.StatusBadRequest, api_errors.ErrEmptyCredentials.Error())
// 		return
// 	}

	data, err := h.services.Users.SignIn(credentials)

	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, data)
	return
}

// @Summary User sign up
// @Tags users
// @Description endpoint for user sign up
// @Produce json
// @Param email formData string true "User's email"
// @Param password formData string true "User's password"
// @Success 200 {object} ApiResponse{result=string}
// @Failure 400 {object} ApiResponse{result=string}
// @Failure 500 {object} ApiResponse{result=string}
// @Failure default {object} ApiResponse{result=string}
// @Router /users/sign-up [post]
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	var credentials input.UserCredentials
    	var decoder = schema.NewDecoder()

    	err := decoder.Decode(&credentials, r.PostForm)

    	if err != nil {
    	    jsonResponse(w, http.StatusInternalServerError, err.Error())
        	return
    	}

// 	if email == "" || password == "" {
// 		jsonResponse(w, http.StatusBadRequest, api_errors.ErrEmptyCredentials.Error())
// 		return
// 	}

	err = h.services.Users.SignUp(credentials)

	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, "User was created successfully")
	return

}
