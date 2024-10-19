package main
import (
    "fmt"
    "os"
    "context"
    "github.com/VadimSmD/KR/internal/repo"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/doug-martin/goqu/v9"
    _ "github.com/doug-martin/goqu/v9/dialect/postgres"
)
type UserRepo struct {
    pg *pgxpool.Pool
}
func NewUserRepo( pg *pgxpool.Pool) UserRepo { return &UserRepo{pg} }
var userTable = goqu.T("users")
func (r *UserRepo) Insert(ctx context.Context, user entity.User) (entity.User, error) {
    sql, args, err := r.Builder.Insert(userTable).Prepared(true).Rows(user).Returning(user).ToSql()
    if err != nil {
        return user, fmt.Errorf("Repo - users - builder: %w", err)
    }
    rows, err = r.Pool.Query(ctx, sql, args...)
    if err != nil {
        return (user, fmt.Errorf("Repo -user - exec: %w", err))
    }
    response, err := pgx.CollectOneRow(rows, pgx.CollectByName[entity.user])
    if err != nil {return user, fmt.Errorf("user.go - newuserrepo - collect: %w", err)}
    return response, nil
}
