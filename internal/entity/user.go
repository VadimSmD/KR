package entity

import (
	"time"
)

type User struct {
	Id         int64     `db:"id" goqu:"skipinsert"`
	Name       string    `db:"name"`
	Surname    string    `db:"surname"`
	Nickname   string    `db:"nickname"`
	Date       time.Time `goqu:"skipinsert"`
	HashedPass string    `db:"hashed_pass"`
	Status     *string   `db:"status"`
}
