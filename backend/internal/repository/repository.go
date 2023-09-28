package repository

import (
	"context"

	"backend/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetMonthlyTaxes(ctx context.Context, dni string) (string, error)
	Logout(ctx context.Context, token string) (string, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
