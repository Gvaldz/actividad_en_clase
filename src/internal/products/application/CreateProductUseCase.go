package application

import (
	"fmt"
	"recu/src/internal/products/domain"
)

type CreateProduct struct {
	repo domain.ProductRepository
}

func NewCreateProduct(repo domain.ProductRepository) *CreateProduct {
	return &CreateProduct{repo: repo}
}

func (c *CreateProduct) Execute(product domain.Product) error {
	fmt.Printf("Guardando producto: %+v\n", product)
	return c.repo.CreateProduct(product)
}
