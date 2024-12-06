package controller

import (
	"database/sql"

	"github.com/HerbertCJ/my-store/model"
	repository "github.com/HerbertCJ/my-store/repository/book"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	Db *sql.DB
}

func NewBookController(db *sql.DB) BookInterface {
	return &BookController{
		Db: db,
	}
}

func (b *BookController) Delete(c *gin.Context) {
	DB := b.Db

	var uri model.BookUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": err})
		return
	}

	repository := repository.NewBookRepository(DB)
	delete := repository.Delete(uri.Id)

	if delete {
		c.JSON(200, gin.H{"status": "success", "message": "Book deleted"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "message": "Failed to delete book"})
	}
}

func (b *BookController) GetAll(c *gin.Context) {
	DB := b.Db

	repository := repository.NewBookRepository(DB)
	get := repository.GetAll()

	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get})
	} else {
		c.JSON(500, gin.H{"status": "failed", "message": "Failed to get books"})
	}
}

func (b *BookController) GetById(c *gin.Context) {
	DB := b.Db

	var uri model.BookUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": err})
		return
	}

	repository := repository.NewBookRepository(DB)
	get := repository.GetById(uri.Id)

	if (get != model.Book{}) {
		c.JSON(200, gin.H{"status": "success", "data": get})
	} else {
		c.JSON(500, gin.H{"status": "failed", "message": "Failed to get book"})
	}
}

func (b *BookController) Create(c *gin.Context) {
	DB := b.Db

	var post model.PostBook
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": err})
		return
	}

	repository := repository.NewBookRepository(DB)
	insert := repository.Create(post)

	if insert {
		c.JSON(200, gin.H{"status": "success", "message": "Book created"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "message": "Failed to create book"})
	}
}

func (b *BookController) Update(c *gin.Context) {
	DB := b.Db

	var post model.PostBook
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": err})
		return
	}

	var uri model.BookUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": err})
		return
	}

	repository := repository.NewBookRepository(DB)
	update := repository.Update(uri.Id, post)

	if (update != model.Book{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "message": "Book updated"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "message": "Failed to update book"})
	}
}
