package service

import (
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

func (s *UsersService) SignUp(email string, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	return s.repo.Create(email, string(passwordHash))
}

func (s *UsersService) SignIn(email string, password string) (*Token, error) {
	user, err := s.repo.GetByEmail(email)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	token, err := s.tokenManager.CreateJWT(string(user.Id))

	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: token,
	}, nil
}
