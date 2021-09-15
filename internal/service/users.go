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
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(passwordHash)
	return s.repo.Create(user)
}

func (s *UsersService) SignIn(user *model.User) (*Token, error) {
	user, err := s.repo.GetByEmail(user.Email)

	if err != nil {
		return nil, api_errors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
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
