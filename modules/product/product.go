package product

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	productsGroup := r.Group("/products")
	{
		productsGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Products route",
			})
		})
	}
}
