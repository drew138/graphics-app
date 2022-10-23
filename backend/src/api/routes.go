package api

import (
	"graphics-app-backend/src/middleware"

	"github.com/drew138/go-graphics/filters/kernels"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	home := router.Group("/")
	home.Use(middleware.AuthMiddleware())
	{
		// Get
		router.GET("/image/:id", getImage)
		router.GET("/image", getImages)
		// Delete
		router.DELETE("/image/:id", deleteImage)
		// Post
		image := home.Group("image")
		image.Use(middleware.ImageParsingMiddleware())
		{
			image.POST("/custom", createCustomFilterImage)
			image.POST("/sharpen", createFilterHandler(kernels.Sharpen))
			image.POST("/edge-detection", createFilterHandler(kernels.EdgeDetection))
			image.POST("/gaussian-blur", createFilterHandler(kernels.GaussianBlur))
			image.POST("/box-blur", createFilterHandler(kernels.BoxBlur))
			image.POST("/negative", createNegativeFilterImage)
		}
	}

}
