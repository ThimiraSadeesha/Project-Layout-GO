package route

import (
	"GoPorject/modules/product"
	"GoPorject/modules/shop"
	"GoPorject/modules/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	product.SetupRoutes(r)
	shop.SetupRoutes(r)
	user.SetupRoutes(r)
}
