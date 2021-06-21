package model

type User struct {
	Id                    int    `json:"id"`
	Email                 string `json:"email"`
	Password              string `json:"password,omitempty"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt string `json:"refresh_token_expires_at"`
}
