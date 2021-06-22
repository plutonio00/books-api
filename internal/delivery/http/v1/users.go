package v1

import (
	"github.com/gorilla/mux"
	api_errors "github.com/plutonio00/books-api/internal/error"
	"net/http"
)

func (h *Handler) initUsersRoutes(router *mux.Router) {
	usersRouter := router.PathPrefix(h.baseEndpoint + "users").Subrouter()

	usersRouter.HandleFunc("/sign-up", h.SignUp).Methods("POST")
	usersRouter.HandleFunc("/sign-in", h.SignIn).Methods("POST")
}

// @Summary User signup
// @Tags users-auth
// @Produce json
// @Param email path string true "dvfsfsf"
// @Param password path string true "dvfsfsf"
// @Success 200 {object} service.Token
// @Failure 400,404 {json} json
// @Failure 500 {json} json
// @Failure default {json} json
// @Router /users/sign-in [post]
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if email == "" || password == "" {
		jsonResponse(w, http.StatusBadRequest, api_errors.ErrInvalidCredentials.Error())
		return
	}

	data, err := h.services.Users.SignIn(email, password)

	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, data)
	return
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if email == "" || password == "" {
		jsonResponse(w, http.StatusBadRequest, api_errors.ErrInvalidCredentials.Error())
		return
	}

	err := h.services.Users.SignUp(email, password)

	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, "User was created successfully")
	return

}
