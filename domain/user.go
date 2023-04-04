package domain

import (
	"context"
	"time"

	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/domain/dto"
)

type LoggerInUserData struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type JWTToken struct {
	ExpiredIn time.Duration
	MaxAge    int
	Secret    string
	Message   string
	User      *LoggerInUserData
}

type User struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
	CreatedAt       time.Time `json:"created_at"`
}

type AuthRepository interface {
	User(ctx context.Context, ctr *User) (*User, error)
	SignIn(ctx context.Context, ctr *dto.SignIn) (*JWTToken, error)
}

type AuthUseCase interface {
	User(ctx context.Context, ctr *User) (*User, error)
	SignIn(ctx context.Context, ctr *dto.SignIn) (*JWTToken, error)
}
