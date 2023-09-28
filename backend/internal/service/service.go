package service

import (
	"context"

	"backend/internal/models"
	"backend/internal/repository"
)

// Service is the business logic of the application.
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	GetMonthlyTaxes(ctx context.Context, dni string) (string, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
