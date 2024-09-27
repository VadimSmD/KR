package main
import (
    "fmt"
    "os"
    "context"
    "../entity/user"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/doug-martin/goqu/v9"
    _ "github.com/doug-martin/goqu/v9/dialect/postgres"
)
type UserRepo struct {
    pg *pgxpool.Pool
}
func NewUserRepo( pg *pgxpool.Pool) UserRepo { return &UserRepo{pg} }

func (r *UserRepo) Insert(ctx context.Context, user entity.User) (entity.User, error) {
    sql, args, err := r.Builder.Insert("users").Colums("user_id name surname hashed_pass nickname date_reg status").Values(user.id, user.name, user.surname, user.hashed_pass, user.nickname, user.date, user.status).ToSql()
    if err != nil {
        return (user, fmt.Errorf("Repo - users - builder: %w", err))
    }
    _, err = r.Pool.Exec(ctx, sql, args)
    if err != nil {
        return (user, fmt.Errorf("Repo -user - exec: %w", err))
    }
    return (user, nil)
}
