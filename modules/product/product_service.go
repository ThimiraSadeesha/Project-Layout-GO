package product

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Service interface {
	GetProducts() []Product
}

type serviceImpl struct{}

func NewService() Service {
	return &serviceImpl{}
}

func (s *serviceImpl) GetProducts() []Product {

	return []Product{
		{ID: 1, Name: "Product A"},
		{ID: 2, Name: "Product B"},
	}
}
