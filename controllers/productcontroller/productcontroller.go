package productcontroller

import (
	"go-simple-projects/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"fmt"
)

func Index(c *gin.Context){
	var products []models.Product
	//This is slice and can be used as array, so has a lot of product
	//Learn about struct again
	fmt.Println(products)

	models.DB.Find(&products) //Learn more about pointer again

	//Unmarshal or Decode the body into a map string interface

	// Marshal and Unmarshal are two methods that are used to encode and decode data into different formats.

	// Marshal is used to transform the data into JSON (or any other format) before sending it to another system. Unmarshal is used to transform the data into a Go object after receiving it from another system.

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context){
	var product models.Product
	//While about this, this only one product

	fmt.Println(product)

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Store(c *gin.Context){
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context){
	var product models.Product

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Delete(c *gin.Context){
	// What is *gin.Context?
	// The gin.Context is the most important part of Gin. It carries request details, validates the request, stores the result of validation, and calls the appropriate handler to process the request.
	// What is gin.H?
	// gin.H is a shortcut for map[string]interface{}.

	var product models.Product

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"product": "Product deleted!"})
}
