package main

import (
	"go-simple-projects/models"
	"go-simple-projects/controllers/productcontroller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDB()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product/store", productcontroller.Store) //CRUD
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
