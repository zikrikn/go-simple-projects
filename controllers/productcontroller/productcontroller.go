package productcontroller

import (
	"go-simple-projects/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context){
	var products []models.Product
	//Learn about struct again

	models.DB.Find(&products) //Learn more about pointer again

	//Unmarshal or Decode the body into a map string interface

	//& * adalah

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context){
	var product models.Product

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
	var product models.Product

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"product": "Product deleted!"})
}
