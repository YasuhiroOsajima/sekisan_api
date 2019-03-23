package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3@v1.9.0"
)

var db *sqlx.DB

func init() {
	db, err = sqlx.Open("sqlite3", "./sekisan.sqlite3")
	if err != nil {
		log.Fatalf("DB connect failed. err: %s", err)
		panic(err)
	}
}

func Db_exec(q string) {
	var _, err = db.Exec(q)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
