package main

import (
	"context"
	"sync"

	userapi "github.com/VadimSmD/KR/internal/repo/userapi"
)

type userService struct {
	users map[int64]userapi.User
	id    int64
	mux   sync.Mutex
}

func (u *userService) CreateUser(ctx context.Context, req CreateUserReq) (*UserStatusCode, error) {

}
func (u *userService) CreateUsersWithListInput(ctx context.Context, req []User) (CreateUsersWithListInputRes, error) {
}

func (u *userService) DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error) {
}

func (u *userService) GetUserByName(ctx context.Context, params GetUserByNameParams) (GetUserByNameRes, error) {
}

func (u *userService) LoginUser(ctx context.Context, params LoginUserParams) (LoginUserRes, error) {}

func (u *userService) LogoutUser(ctx context.Context) (*LogoutUserDef, error) {}

func (u *userService) UpdateUser(ctx context.Context, req UpdateUserReq, params UpdateUserParams) (*UpdateUserDef, error) {
}
