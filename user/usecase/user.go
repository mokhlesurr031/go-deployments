package usecase

import (
	"context"

	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/domain"
	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/domain/dto"
)

// New return new usecase for user
func New(repo domain.AuthRepository) domain.AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

type AuthUseCase struct {
	repo domain.AuthRepository
}

func (a *AuthUseCase) User(ctx context.Context, ctr *domain.User) (*domain.User, error) {
	return a.repo.User(ctx, ctr)
}

func (a *AuthUseCase) SignIn(ctx context.Context, ctr *dto.SignIn) (*domain.JWTToken, error) {
	return a.repo.SignIn(ctx, ctr)
}
