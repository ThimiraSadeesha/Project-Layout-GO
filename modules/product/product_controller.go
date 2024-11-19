package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) GetProducts(c *gin.Context) {
	products := ctrl.service.GetProducts()
	c.JSON(http.StatusOK, gin.H{
		"message": "Products fetched successfully",
		"data":    products,
	})
}
func (ctrl *Controller) ViewProduct(c *gin.Context) {
	products := ctrl.service.GetProducts()
	c.JSON(http.StatusOK, gin.H{
		"message": "Products fetched successfully",
		"data":    products,
	})
}

func (ctrl *Controller) SetupRoutes(r *gin.Engine) {
	productsGroup := r.Group("/products")
	{
		productsGroup.GET("/", ctrl.GetProducts)
		productsGroup.GET("view/", ctrl.ViewProduct)
	}
}
