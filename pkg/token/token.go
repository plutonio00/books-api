package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

type TokenManager struct {
	signingKey string
}

func NewTokenManager(signingKey string) (*TokenManager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &TokenManager{signingKey: signingKey}, nil
}

func (t *TokenManager) CreateJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject: userId,
	})

	return token.SignedString([]byte(t.signingKey))
}

func (t *TokenManager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
    		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
    			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    		}

    		return []byte(t.signingKey), nil
    	})
    	if err != nil {
    		return "", err
    	}

    	claims, ok := token.Claims.(jwt.MapClaims)
    	if !ok {
    		return "", fmt.Errorf("error get user claims from token")
    	}

    	return claims["sub"].(string), nil
}
