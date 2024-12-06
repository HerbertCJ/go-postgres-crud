package app

import (
	"database/sql"

	controller "github.com/HerbertCJ/my-store/controller/book"
	"github.com/HerbertCJ/my-store/db"
	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func (a *App) CreateConeection() {
	db := db.Connectdb()
	a.DB = db
}

func (a *App) Routes() {
	r := gin.Default()
	controller := controller.NewBookController(a.DB)

	r.POST("/books", controller.Create)
	r.GET("/books", controller.GetAll)
	r.GET("/books/:id", controller.GetById)
	r.PUT("/books/:id", controller.Update)
	r.DELETE("/books/:id", controller.Delete)
	a.Router = r
}

func (a *App) Run() {
	a.Router.Run(":8080")
}
