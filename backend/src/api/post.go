package api

import (
	"encoding/json"
	"graphics-app-backend/src/database"
	"graphics-app-backend/src/helpers"
	"image"
	"net/http"

	"github.com/drew138/go-graphics/filters"
	"github.com/drew138/go-graphics/filters/kernels"
	"github.com/gin-gonic/gin"
)

func createCustomFilterImage(c *gin.Context) {
	userId := c.GetString("userId")
	var kernel kernels.Kernel
	json.Unmarshal([]byte(c.Request.FormValue("kernel")), &kernel)
	if len(kernel) != 3 || len(kernel[0]) != 3 || len(kernel[1]) != 3 || len(kernel[2]) != 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Kernel"})
		return
	}

	decodedImage, ok := c.Get("image")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	outputImage := filters.ApplyFilter(decodedImage.(image.Image), kernel)

	url, err := helpers.UploadImage(outputImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newImage, err := database.InsertImage(userId, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newImage)
}

func createFilterHandler(kernel kernels.Kernel) func(*gin.Context) {
	return func(c *gin.Context) {
		userId := c.GetString("userId")
		decodedImage, ok := c.Get("image")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}
		outputImage := filters.ApplyFilter(decodedImage.(image.Image), kernel)

		url, err := helpers.UploadImage(outputImage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		newImage, err := database.InsertImage(userId, url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, newImage)
	}
}

func createNegativeFilterImage(c *gin.Context) {
	userId := c.GetString("userId")
	decodedImage, ok := c.Get("image")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	outputImage := filters.CreateNegativeImage(decodedImage.(image.Image))

	url, err := helpers.UploadImage(outputImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newImage, err := database.InsertImage(userId, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newImage)
}
