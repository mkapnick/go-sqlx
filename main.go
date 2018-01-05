package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Build struct {
	Id         int        `db:"id"`
	UserId     string     `db:"userId"`
	UserName   string     `db:"userName"`
	BranchName string     `db:"branchName"`
	SlackBody  string     `db:"slackBody"`
	CreatedAt  *time.Time `db:"createdAt"`
	UpdatedAt  *time.Time `db:"updatedAt"`
}

func main() {
	db := sqlx.MustConnect("mysql", "root@tcp(127.0.0.1:3306)/jack?charset=utf8&parseTime=True")
	rows, err := db.Queryx("SELECT userId, userName, branchName from build")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var b Build
		err = rows.StructScan(&b)

		if err != nil {
			panic(err)
		}

		println(b.UserId)
		println(b.UserName)
		println(b.BranchName)
	}
}
