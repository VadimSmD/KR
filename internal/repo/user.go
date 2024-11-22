package repo

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/VadimSmD/KR/internal/entity"
)

type UserRepo struct {
	pg *pgxpool.Pool
}

type PartialIdent struct {
	userId   int
	nickname string
	hashedp  string
}

func NewUserRepo(pg *pgxpool.Pool) *UserRepo {
	return &UserRepo{pg: pg}
}

var userTable = goqu.T("users")

func (r *UserRepo) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	sql, args, err := goqu.Insert("users").Prepared(true).Rows(user).Returning(goqu.T("users").All).ToSQL()
	if err != nil {
		return user, fmt.Errorf("Repo - users - insert: %w", err)
	}
	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return user, fmt.Errorf("Repo -user - insert_exec: %w", err)
	}
	response, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return user, fmt.Errorf("user.go - newuserrepo - collect: %w", err)
	}
	return response, nil
}

func (r *UserRepo) Delete(ctx context.Context, userId int) error {
	sql, args, err := goqu.Delete("users").Where(goqu.Ex{"id": userId}).ToSQL()
	if err != nil {
		return fmt.Errorf("Repo - users - delete: %w", err)
	}
	if _, err = r.pg.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("Repo - users - delete: %w", err)
	}
	return nil
}

func (r *UserRepo) Edit(ctx context.Context, user entity.User) (entity.User, error) {
	sql, args, err := goqu.Update("users").Where(goqu.Ex{"Id": user.Id}).Set(user).ToSQL()
	if err != nil {
		return user, fmt.Errorf("Repo - users - update: %w", err)
	}
	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return user, fmt.Errorf("Repo -user - update_exec: %w", err)
	}
	response, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return user, fmt.Errorf("user.go - update - collect: %w", err)
	}
	return response, nil
}

func (r *UserRepo) Select(ctx context.Context, ident PartialIdent) (entity.User, error) {
	ds := goqu.From("users").Where(
		goqu.And(
			goqu.C("login").Eq(ident.nickname),
			goqu.C("hashed_pass").Eq(ident.hashedp),
			goqu.C("id").Eq(ident.userId),
		),
	)
	sql, args, err := ds.Prepared(true).ToSQL()
	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return entity.User{}, fmt.Errorf("Repo -user - select_exec: %w", err)
	}
	response, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return entity.User{}, fmt.Errorf("user.go - select - collect: %w", err)
	}
	return response, nil

}
