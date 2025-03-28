package infrastructure

import (
	"recu/src/internal/products/application"
	"recu/src/internal/products/infrastructure/controllers"
)

type ProductDependences struct {
	repo *MemoryProductRepository
}

func NewProductDependences() *ProductDependences {
	repo := NewMemoryProductRepository()
	return &ProductDependences{repo: repo}
}

func (d *ProductDependences) GetRoutes() *ProductRoutes {
	createProductUseCase := application.NewCreateProduct(d.repo)
	getNewProductsCountUseCase := application.NewGetNewProductsCount(d.repo)
	getDiscountProductsCountUseCase := application.NewGetDiscountProductsCount(d.repo)

	createProductController := controllers.NewCreateProductController(createProductUseCase)
	getNewProductsCountController := controllers.NewGetNewProductsCountController(getNewProductsCountUseCase)
	getDiscountProductsCountController := controllers.NewGetDiscountProductsCountController(getDiscountProductsCountUseCase)

	return NewProductRoutes(
		createProductController,
		getNewProductsCountController,
		getDiscountProductsCountController,
	)
}