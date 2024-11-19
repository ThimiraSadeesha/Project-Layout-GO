package main

import (
	"GoPorject/modules/product"
	"GoPorject/modules/shop"
	"GoPorject/modules/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	user.SetupRoutes(r)
	product.SetupRoutes(r)
	shop.SetupRoutes(r)
	r.Run(":8080")
}
