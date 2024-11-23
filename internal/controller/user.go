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
		Delete(ctx context.Context, deleteOpts string) error
		Edit(ctx context.Context, user entity.User) (entity.User, error)
		Select(ctx context.Context, selectOpts entity.PartialIdent) (entity.User, error)
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

func (c *UserController) DeleteUser(ctx context.Context, params oas.DeleteUserParams) (oas.DeleteUserRes, error) {
	err := c.service.Delete(ctx, params.Username)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *UserController) GetUserByName(ctx context.Context, params oas.GetUserByNameParams) (oas.GetUserByNameRes, error) {
	selectOpts := entity.PartialIdent{
		Nickname: params.Username,
		UserId:   0,
		Hashedp:  "",
	}
	response, err := c.service.Select(ctx, selectOpts)
	if err != nil {
		return nil, err
	}
	return userToAPI(response), nil
}

func (c *UserController) UpdateUser(ctx context.Context, req oas.CreateUserRequest, params oas.UpdateUserParams) (entity.User, error) {
	updateOpts := entity.User{
		Id:         0,
		Name:       req.Name,
		Surname:    req.Surname,
		Nickname:   req.Nickname,
		HashedPass: req.HashedPass,
	}
	resp, err := c.service.Edit(ctx, updateOpts)
	if err != nil {
		return entity.User{}, err
	}
	return resp, nil
}

func (c *UserController) CreateUsersWithListInput(ctx context.Context, req []oas.User) error {
	for _, user := range req {
		update := oas.CreateUserRequest{
			Nickname:   user.Nickname,
			Name:       user.Name,
			Surname:    user.Surname,
			HashedPass: user.HashedPass.Value,
		}
		_, err := c.CreateUser(ctx, &update)
		if err != nil {
			return err
		}
	}
	return nil
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

func (c *UserController) LoginUser(ctx context.Context, params oas.LoginUserParams) (oas.LoginUserRes, error) {
	return nil, fmt.Errorf("implement me")
}

func (c *UserController) LogoutUser(ctx context.Context) error {
	return fmt.Errorf("implement me")
}
