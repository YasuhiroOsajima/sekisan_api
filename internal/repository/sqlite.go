package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open("sqlite3", "./sekisan.sqlite3")
	if err != nil {
		log.Fatalf("DB connect failed. err: %s", err)
		panic(err)
	}
}
