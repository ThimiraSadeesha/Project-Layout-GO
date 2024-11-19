package product

import (
	"github.com/gin-gonic/gin"
)

type Module struct {
	Controller *Controller
}

func NewModule() *Module {
	productService := NewService()
	productController := NewController(productService)

	return &Module{
		Controller: productController,
	}
}

func (m *Module) RegisterRoutes(r *gin.Engine) {
	m.Controller.SetupRoutes(r)
}
