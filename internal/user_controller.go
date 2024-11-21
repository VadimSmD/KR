package controller

import (
	"context"
	"sync"

	ogen "github.com/VadimSmD/KR/internal/repo/userapi"
	logic "github.com/VadimSmD/KR/internal/repo/user.go"
)

type UserController struct {
	ogen.UserHandler
	userService UserService
}

func (c *UserController) CreateUser(ctx context.Context, req CreateUserReq) (*UserStatusCode, error) {
	user := User{Name : req.Name, Surname : req.Surname, Nickname : req.Nickname, Date : req.Date, HashedPass : req.HashedPass, Status : req.Status}
	return logic.Insert(ctx, user)
}
func (c *UserController) CreateUsersWithListInput(ctx context.Context, req []User) (CreateUsersWithListInputRes, error) {
	for i:=0; i<len(req); i++ {
		logic.Insert(ctx, req[i])
	}
}

func (c *UserController) DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error) {
	return logic.Delete(ctx, params)
}

func (c UserController) GetUserByName(ctx context.Context, params GetUserByNameParams) (GetUserByNameRes, error) {
	return logic.Select(ctx, params)
}

func (c UserController) LoginUser(ctx context.Context, params LoginUserParams) (LoginUserRes, error) {
}

func (c UserController) LogoutUser(ctx context.Context) (*LogoutUserDef, error) {}

func (c UserController)) UpdateUser(ctx context.Context, req UpdateUserReq, params UpdateUserParams) (*UpdateUserDef, error) {
	user := User{Id : req.Id , Name : params.Name, Surname : params.Surname, Nickname : params.Nickname, Date : params.Date, HashedPass : params.HashedPass, Status : params.Status}
	return logic.Edit(ctx, user)
}
