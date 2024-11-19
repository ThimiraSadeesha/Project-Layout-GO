package user

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	usersGroup := r.Group("/users")
	{
		usersGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Users route",
			})
		})
	}
}
