package application

import "recu/src/internal/products/domain"

type GetNewProductsCount struct {
	repo domain.ProductRepository
}

func NewGetNewProductsCount(repo domain.ProductRepository) *GetNewProductsCount {
	return &GetNewProductsCount{repo: repo}
}

func (uc *GetNewProductsCount) Execute() int {
	return uc.repo.GetNewProductsCount()
}