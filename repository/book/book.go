package repository

import (
	"database/sql"
	"log"

	"github.com/HerbertCJ/my-store/model"
)

type BookRepository struct {
	Db *sql.DB
}

func NewBookRepository(db *sql.DB) BookInterface {
	return &BookRepository{
		Db: db,
	}
}

func (b *BookRepository) Delete(id uint) bool {
	_, err := b.Db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (b *BookRepository) GetAll() []model.Book {
	query, err := b.Db.Query("SELECT * FROM books")

	if err != nil {
		log.Println(err)
		return nil
	}

	var books []model.Book
	if query != nil {
		defer query.Close()
		for query.Next() {
			var (
				id     uint
				title  string
				author string
			)

			err := query.Scan(&id, &title, &author)

			if err != nil {
				log.Println(err)
			}

			book := model.Book{Id: id, Title: title, Author: author}
			books = append(books, book)
		}
	}

	return books
}

func (b *BookRepository) GetById(id uint) model.Book {
	query, err := b.Db.Query("SELECT * FROM books WHERE id=$1", id)

	if err != nil {
		log.Println(err)
		return model.Book{}
	}

	var book model.Book
	if query != nil {
		defer query.Close()
		for query.Next() {
			var (
				id     uint
				title  string
				author string
			)

			err := query.Scan(&id, &title, &author)
			if err != nil {
				log.Println(err)
			}

			book = model.Book{Id: id, Title: title, Author: author}
		}
	}

	return book
}

func (b *BookRepository) Create(book model.PostBook) bool {
	stmt, err := b.Db.Prepare("INSERT INTO books (title, author) VALUES ($1, $2)")

	if err != nil {
		log.Println(err)
		return false
	}

	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Author)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (b *BookRepository) Update(id uint, book model.PostBook) model.Book {
	_, err := b.Db.Exec("UPDATE books SET title=$1, author=$2 WHERE id=$3", book.Title, book.Author, id)

	if err != nil {
		log.Println(err)
		return model.Book{}
	}

	return b.GetById(id)
}
