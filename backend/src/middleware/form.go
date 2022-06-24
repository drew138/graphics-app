package middleware

import (
	"image"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ImageParsingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		file, err := fileHeader.Open()

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		defer file.Close()
		decodedImage, format, err := image.Decode(file)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		c.Set("image", decodedImage)
		c.Set("format", format)
		c.Next()
	}
}
