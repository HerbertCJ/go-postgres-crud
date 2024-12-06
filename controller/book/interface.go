package controller

import "github.com/gin-gonic/gin"

type BookInterface interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
