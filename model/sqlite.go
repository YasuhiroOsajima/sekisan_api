package model

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type QueryExecutor interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
}

func Db_connect() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./sekisan.sqlite3")
	if err != nil {
		log.Fatalf("DB connect failed. err: %s", err)
		panic(err)
	}

	return db
}

func Db_exec(db *sql.DB, q string) {
	var _, err = db.Exec(q)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
