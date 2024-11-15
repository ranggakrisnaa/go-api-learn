package user

import (
	"context"
	"go-api-learn/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	Store(context.Context, *domain.User) (*uuid.UUID, error)
}

type Service struct {
	userRepo UserRepository
}

func NewService(u UserRepository) *Service {
	return &Service{
		userRepo: u,
	}
}

func (s *Service) Store(context.Context, *domain.User) (*uuid.UUID, error) {
	return nil, nil
}
