package controller

import (
	"context"
	"fmt"

	"github.com/VadimSmD/KR/gen/oas"
	"github.com/VadimSmD/KR/internal/entity"
)

type (
	UserController struct {
		oas.UserHandler
		service UserServiceProvider
	}
	UserServiceProvider interface {
		Create(ctx context.Context, user entity.User) (entity.User, error)
	}
)

func NewUserController(service UserServiceProvider) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) CreateUser(ctx context.Context, req *oas.CreateUserRequest) (*oas.User, error) {
	createOpts := entity.User{
		Id:         0,
		Name:       req.Name,
		Surname:    req.Surname,
		Nickname:   req.Nickname,
		HashedPass: req.HashedPass,
	}
	response, err := c.service.Create(ctx, createOpts)
	if err != nil {
		return nil, err
	}
	return userToAPI(response), nil
}

func userToAPI(user entity.User) *oas.User {
	return &oas.User{
		ID:       user.Id,
		Name:     user.Name,
		Surname:  user.Surname,
		Nickname: user.Nickname,
		Date:     user.Date,
	}
}

func (c *UserController) CreateUsersWithListInput(ctx context.Context, req []oas.User) (*oas.User, error) {
	return nil, fmt.Errorf("Implement me")
}

func (c *UserController) DeleteUser(ctx context.Context, params oas.DeleteUserParams) (oas.DeleteUserRes, error) {
	return nil, fmt.Errorf("Implement me")
}

func (c *UserController) GetUserByName(ctx context.Context, params oas.GetUserByNameParams) (oas.GetUserByNameRes, error) {
	return nil, fmt.Errorf("Implement me")
}

func (c *UserController) LoginUser(ctx context.Context, params oas.LoginUserParams) (oas.LoginUserRes, error) {
	return nil, fmt.Errorf("Implement me")
}

func (c *UserController) LogoutUser(ctx context.Context) error {
	return fmt.Errorf("Implement me")
}

func (c *UserController) UpdateUser(ctx context.Context, req oas.UpdateUserReq, params oas.UpdateUserParams) error {
	return fmt.Errorf("Implement me")
}
