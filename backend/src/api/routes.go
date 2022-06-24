package api

import (
	"github.com/drew138/go-graphics/filters/kernels"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Get
	router.GET("/image/:id", getImage)
	router.GET("/image", getImages)
	// Post
	router.POST("/image/custom", createCustomFilterImage)
	router.POST("/image/sharpen", createFilterHandler(kernels.Sharpen))
	router.POST("/image/edge-detection", createFilterHandler(kernels.EdgeDetection))
	router.POST("/image/gaussian-blur", createFilterHandler(kernels.GaussianBlur))
	router.POST("/image/box-blur", createFilterHandler(kernels.BoxBlur))
	router.POST("/image/negative", createNegativeFilterImage)
	// Delete
	router.DELETE("/image/:id", deleteImage)

}
