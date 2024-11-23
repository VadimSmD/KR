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
		Delete(ctx context.Context, nickname string) error
		Edit(ctx context.Context, user entity.User) (entity.User, error)
		Select(ctx context.Context, ident entity.PartialIdent) (entity.User, error)
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

func (s *UserService) Delete(ctx context.Context, deleteOpts string) error {
	err := s.repo.Delete(ctx, deleteOpts)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

func (s *UserService) Edit(ctx context.Context, user entity.User) (entity.User, error) {
	resp, err := s.repo.Edit(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("edit: %w", err)
	}
	return resp, nil
}

func (s *UserService) Select(ctx context.Context, selectOpts entity.PartialIdent) (entity.User, error) {
	resp, err := s.repo.Select(ctx, selectOpts)
	if err != nil {
		return entity.User{}, fmt.Errorf("select: %w", err)
	}
	return resp, nil
}
