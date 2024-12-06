package repository

import "github.com/HerbertCJ/my-store/model"

type BookInterface interface {
	GetAll() []model.Book
	GetById(id uint) model.Book
	Create(book model.PostBook) bool
	Update(id uint, book model.PostBook) model.Book
	Delete(id uint) bool
}
