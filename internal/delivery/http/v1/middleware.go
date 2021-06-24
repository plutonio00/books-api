package v1

import (
	"context"
	// 	"errors"
	api_errors "github.com/plutonio00/books-api/internal/error"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) authWithToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := h.parseAuthHeader(w, r)

		if err != nil {
			jsonResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		id, err := h.tokenManager.Parse(token)

		if err != nil {
			jsonResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "user", id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) parseAuthHeader(w http.ResponseWriter, r *http.Request) (string, error) {
	header := r.Header.Get(authorizationHeader)
	if header == "" {
		return "", api_errors.ErrUnauthorized
		// 		return "", errors.New("Empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", api_errors.ErrUnauthorized
		// 		return "", errors.New("Invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", api_errors.ErrUnauthorized
		// 		return "", errors.New("Token is empty")
	}

	return headerParts[1], nil
}
