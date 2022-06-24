package api

import (
	"graphics-app-backend/src/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deleteImage(c *gin.Context) {
	userId := c.GetInt("userId")
	imageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DeleteImage(imageId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
