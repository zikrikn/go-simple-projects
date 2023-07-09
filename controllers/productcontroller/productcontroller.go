package productcontroller

import (
	"encoding/json"
	"go-simple-projects/models"
	"net/http"
	// "strconv"

	// "fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var products []models.Product
	//This is slice and can be used as array, so has a lot of product

	models.DB.Find(&products) //Learn more about pointer again
	// Penjelasa tentang models.DB.Find(&products)? models.DB.Find(&products) is used to retrieve all records from the products table and store them in the products variable.

	//Unmarshal or Decode the body into a map string interface

	// Marshal and Unmarshal are two methods that are used to encode and decode data into different formats.

	// Marshal is used to transform the data into JSON (or any other format) before sending it to another system. Unmarshal is used to transform the data into a Go object after receiving it from another system.

	c.JSON(http.StatusOK, gin.H{"products": products})

	//g.H is a shortcut for map[string]interface{}
}

func Show(c *gin.Context){
	var product models.Product
	//While about this, this only one product

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&product).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
			// digunakan untuk menghentikan eksekusi lebih lanjut dan langsung mengirimkan respons JSON dengan status HTTP 400 Bad Request
			// c.AbortWithStatusJSON function is used to abort a request with a custom error message.
			// batalkan dengan status json
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
	// c.JSON fungsinya untuk mengembalikan data dalam bentuk JSON
}

func Store(c *gin.Context){
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} // err := c.ShouldBindJSON(&product); err != nil, ini tuh kaya for, soalnya assign value ke product dan kondisinya.
	// Kalau nill is a good thing, kalau tidak nill, maka akan return error karena tidak sesuai dengan struct data yang dibuat
	// c itu melempar data-nya, isinya body-nya

	// c.ShouldBindJSON function is used to bind the request body to a struct. If the request body doesn’t match the struct, it will return an error.

	models.DB.Create(&product) // Jadi isi dari &product itu adalah struct? Yes, &product is a struct, jadi GORM itu mengambil data dari struct dan mengirimkan ke database
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context){
	var product models.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated!"})
}

func Delete(c *gin.Context) {
	// What is *gin.Context?
	// The gin.Context is the most important part of Gin. It carries request details, validates the request, stores the result of validation, and calls the appropriate handler to process the request. // We have to this more details
	// What is gin.H?
	// gin.H is a shortcut for map[string]interface{}.

	var product models.Product

	var input struct {
		ID json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := input.ID.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": "Product deleted!"})
}
