package service

import (
	"context"
	"fmt"

	"github.com/VadimSmD/KR/internal/entity"
)

type (
	UserService struct {
		repo UserRepoProvider
	}
	UserRepoProvider interface {
		Insert(ctx context.Context, user entity.User) (entity.User, error)
	}
)

func NewUserService(repo UserRepoProvider) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user entity.User) (entity.User, error) {
	resp, err := s.repo.Insert(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("create: %w", err)
	}
	return resp, nil
}
