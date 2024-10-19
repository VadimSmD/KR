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
        return user, fmt.Errorf("Repo - users - insert: %w", err)
    }
    rows, err = r.Pool.Query(ctx, sql, args...)
    if err != nil {
        return (user, fmt.Errorf("Repo -user - insert_exec: %w", err))
    }
    response, err := pgx.CollectOneRow(rows, pgx.CollectByName[entity.user])
    if err != nil {return user, fmt.Errorf("user.go - newuserrepo - collect: %w", err)}
    return response, nil
}

func (r *UserRepo) Delete(ctx context.Context, user_id int) (error) {
	sql, args, err := r.Builder.Delete(userTable).Where(goqu.Ex{"id":user_id}).ToSql()
    if err != nil {
        return fmt.Errorf("Repo - users - delete: %w", err)
    }
    return nil
}

func (r *UserRepo) Edit(ctx context.Context, user entity.User) (entity.User, error) {
    sql, args, err := r.Builder.Update(userTable).Where(goqu.Ex{"id":user.id}).Set(user).ToSql()
    if err != nil {
        return user, fmt.Errorf("Repo - users - update: %w", err)
    }
    return response, nil
}

func (r *UserRepo) Get(ctx context.Context, user_id int) () {
	sql, args, err := r.Builder.Select(userTable).Where(goqu.Ex{"id":user_id}).Returning(user).ToSql()
    if err != nil {
        return user, fmt.Errorf("Repo - users - get: %w", err)
    }
    rows, err = r.Pool.Query(ctx, sql, args...)
    if err != nil {
        return (user, fmt.Errorf("Repo -user - get_exec: %w", err))
    }
    response, err := pgx.CollectOneRow(rows, pgx.CollectByName[entity.user])
    if err != nil {return user, fmt.Errorf("user.go - get_user - collect: %w", err)}
    return response, nil
}
}