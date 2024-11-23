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

func NewUserRepo(pg *pgxpool.Pool) *UserRepo {
	return &UserRepo{pg: pg}
}

func (r *UserRepo) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	sql, args, err := goqu.Insert("users").Prepared(true).Rows(user).Returning(goqu.T("users").All).ToSQL()
	if err != nil {
		return user, fmt.Errorf("repo - users - insert: %w", err)
	}
	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return user, fmt.Errorf("repo -user - insert_exec: %w", err)
	}
	response, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return user, fmt.Errorf("user.go - newuserrepo - collect: %w", err)
	}
	return response, nil
}

func (r *UserRepo) Delete(ctx context.Context, nickname string) error {
	sql, args, err := goqu.Delete("users").Where(goqu.Ex{"nickname": nickname}).ToSQL()
	if err != nil {
		return fmt.Errorf("repo - users - delete: %w", err)
	}
	if _, err = r.pg.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("repo - users - delete: %w", err)
	}
	return nil
}

func (r *UserRepo) Edit(ctx context.Context, user entity.User) (entity.User, error) {
	sql, args, err := goqu.Update("users").Where(goqu.Ex{"Id": user.Id}).Set(user).ToSQL()
	if err != nil {
		return user, fmt.Errorf("repo - users - update: %w", err)
	}
	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return user, fmt.Errorf("repo -user - update_exec: %w", err)
	}
	response, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return user, fmt.Errorf("user.go - update - collect: %w", err)
	}
	return response, nil
}

func (r *UserRepo) Select(ctx context.Context, ident entity.PartialIdent) (entity.User, error) {
	ds := goqu.From("users")
	whereConditions := goqu.Ex{}
	if ident.Nickname != "" {
		whereConditions["login"] = ident.Nickname
	}
	if ident.Hashedp != "" {
		whereConditions["hashed_pass"] = ident.Hashedp
	}

	if ident.UserId != 0 {
		whereConditions["id"] = ident.UserId
	}
	sql, args, _ := ds.Prepared(true).ToSQL()
	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return entity.User{}, fmt.Errorf("repo -user - select_exec: %w", err)
	}
	response, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[entity.User])
	if err != nil {
		return entity.User{}, fmt.Errorf("user.go - select - collect: %w", err)
	}
	return response, nil

}
