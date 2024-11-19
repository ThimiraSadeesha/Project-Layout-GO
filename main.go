package main

import (
	"GoPorject/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	route.SetupRoutes(r)
	r.Run(":8080")
}
