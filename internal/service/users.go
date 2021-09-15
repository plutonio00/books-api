package service

import (
	api_errors "github.com/plutonio00/books-api/internal/error"
	"github.com/plutonio00/books-api/internal/model"
	"github.com/plutonio00/books-api/internal/repository"
	"github.com/plutonio00/books-api/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	repo         repository.Users
	tokenManager *token.TokenManager
}

func NewUsersService(repo repository.Users, tokenManager *token.TokenManager) *UsersService {

	return &UsersService{
		repo:         repo,
		tokenManager: tokenManager,
	}
}

func (s *UsersService) SignUp(user *model.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	credentials.Password = string(passwordHash)
	return s.repo.Create(credentials)
}

func (s *UsersService) SignIn(user *model.User) (*Token, error) {
	user, err := s.repo.GetByEmail(credentials.Email)

	if err != nil {
		return nil, api_errors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, api_errors.ErrInvalidCredentials
	}

	token, err := s.tokenManager.CreateJWT(string(user.Id))

	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: token,
	}, nil
}
