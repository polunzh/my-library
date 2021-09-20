package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "./data/book.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS book(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	isbn TEXT,
	remark TEXT,
	created_at TEXT,
	updated_at TEXT
	)
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("create table book failed%q: %s\n", err, sqlStmt)
	}
}

type Book struct {
	id         int64
	title      string
	isbn       string
	remark     string
	created_at string
	updated_at string
}

func Insert(data *Book) (int64, error) {
	result, err := db.Exec(fmt.Sprintf(`INSERT INTO
	book(title, isbn, remark, created_at, updated_at)
	VALUES("%s", "%s", "%s", "%s", "%s")
	`, data.title, data.isbn, data.remark, data.created_at, data.updated_at))

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func FindByISBN(isbn string) (*Book, error) {
	rows, err := db.Query(fmt.Sprintf(`SELECT
	id, title, isbn, remark, created_at, updated_at
	FROM book WHERE isbn=%s`, isbn))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var id int64
		var title string
		var isbn string
		var remark string
		var created_at string
		var updated_at string

		err = rows.Scan(&id, &title, &isbn, &remark, &created_at, &updated_at)
		if err != nil {
			return nil, err
		}

		return &Book{id: id, title: title, isbn: isbn, remark: remark, created_at: created_at, updated_at: updated_at}, nil
	}

	return nil, rows.Err()
}
