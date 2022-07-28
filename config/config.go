package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "data.db")
	q := `CREATE TABLE IF NOT EXISTS 'note' (
		'id'	INTEGER PRIMARY KEY AUTOINCREMENT,
		'title'	TEXT,
		'body'	TEXT,
		'color'	INTEGER
	);`
	db.Exec(q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return db, err
	}
}
