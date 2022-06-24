package api

import (
	"graphics-app-backend/src/database"
	"graphics-app-backend/src/helpers"
	"graphics-app-backend/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getImage(c *gin.Context) {
	authId := c.GetString("auth_id")
	imageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	image, err := database.SelectImage(imageId, authId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, image)
}

type body struct {
	Previous string          `json:"previous"`
	Next     string          `json:"next"`
	Results  []*models.Image `json:"results"`
}

func getImages(c *gin.Context) {
	// Get user data
	userId := c.GetString("userId")
	// Get query params
	from, err := strconv.Atoi(c.Query("from"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	to, err := strconv.Atoi(c.Query("to"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Get requested images
	images, err := database.SelectPaginatedUserImages(userId, from, to)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	previousUrl := helpers.FormatPaginatedURL("image", from-25, from)
	nextUrl := helpers.FormatPaginatedURL("image", to, to+25)
	c.JSON(http.StatusCreated, body{Previous: previousUrl, Next: nextUrl, Results: images})
}
