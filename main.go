package main

import (
	"github.com/IMarcellinus/go-restapi-gin/controllers/bookcontroller"
	"github.com/IMarcellinus/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnDB()

	r.GET("/", bookcontroller.Welcome)
	r.GET("/api/book", bookcontroller.Index)
	r.GET("/api/book/:id", bookcontroller.Show)
	r.POST("/api/book", bookcontroller.Create)
	r.PUT("/api/book/:id", bookcontroller.Update)
	r.DELETE("/api/book", bookcontroller.Delete)

	r.Run()
}
