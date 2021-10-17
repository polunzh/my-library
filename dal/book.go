package dal

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := sql.Open("sqlite3", "./data/book.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS book(
	id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
	title TEXT CHECK(LENGTH(title) >= 1) NOT NULL UNIQUE,
	isbn TEXT CHECK(LENGTH(isbn) >= 1) NOT NULL UNIQUE,
	purchase_from TEXT,
	remark TEXT,
	created_at TEXT CHECK(LENGTH(created_at) >= 1),
	updated_at TEXT CHECK(LENGTH(updated_at) >= 1)
	)
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("create table book failed%q: %s\n", err, sqlStmt)
	}
}

type Book struct {
	Id           int64
	Title        string
	Isbn         string
	PurchaseFrom string
	Remark       string
	CreatedAt    string
	UpdatedAt    string
}

func Insert(data *Book) (int64, error) {
	db, err := sql.Open("sqlite3", "./data/book.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	now := time.Now().Format(time.RFC3339)
	data.CreatedAt = now
	data.UpdatedAt = now
	result, err := db.Exec(fmt.Sprintf(`INSERT INTO
	book(title, isbn, remark, created_at, updated_at)
	VALUES("%s", "%s", "%s", "%s", "%s")
	`, data.Title, data.Isbn, data.Remark, data.CreatedAt, data.UpdatedAt))

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func read(rows *sql.Rows) (*Book, error) {
	if rows.Next() {
		var id int64
		var title string
		var isbn string
		var remark string
		var createdAt string
		var updatedAt string

		err := rows.Scan(&id, &title, &isbn, &remark, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		return &Book{Id: id, Title: title, Isbn: isbn, Remark: remark, CreatedAt: createdAt, UpdatedAt: updatedAt}, nil
	}

	return nil, rows.Err()
}

func FindByISBN(isbn string) (*Book, error) {
	db, err := sql.Open("sqlite3", "./data/book.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT
	id, title, isbn, remark, created_at, updated_at
	FROM book WHERE isbn=%s`, isbn))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return read(rows)
}

func FindById(id string) (*Book, error) {
	db, err := sql.Open("sqlite3", "./data/book.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT
	id, title, isbn, remark, created_at, updated_at
	FROM book WHERE id=%s`, id))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return read(rows)
}

func FindAll() ([]*Book, error) {
	db, err := sql.Open("sqlite3", "./data/book.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf(`SELECT
	id, title, isbn, remark, created_at, updated_at
	FROM book`))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*Book

	for rows.Next() {
		var id int64
		var title string
		var isbn string
		var purchaseFrom string
		var remark string
		var createdAt string
		var updatedAt string

		err = rows.Scan(&id, &title, &isbn, &remark, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		books = append(books, &Book{Id: id, Title: title, Isbn: isbn, PurchaseFrom: purchaseFrom, Remark: remark, CreatedAt: createdAt, UpdatedAt: updatedAt})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return books, nil
}
