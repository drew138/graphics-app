package main

import (
	"graphics-app-backend/src/api"
	"graphics-app-backend/src/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(middleware.AuthMiddleware())
	r.Use(gin.Recovery())
	// url := helpers.FormatPaginatedURL("image", 25, 50)
	// fmt.Println(url)
	api.RegisterRoutes(r)
	r.Run()
}
