package shop

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	shopsGroup := r.Group("/shop")
	{
		shopsGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Shop route",
			})
		})
	}
}
