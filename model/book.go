package model

type Book struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type PostBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookUri struct {
	Id uint `uri:"id" binding:"required,number"`
}
