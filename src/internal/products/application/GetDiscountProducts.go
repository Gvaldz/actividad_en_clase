package application

import (
	"recu/src/internal/products/domain"
	"time"
)

type GetDiscountProductsCount struct {
	repo domain.ProductRepository
}

func NewGetDiscountProductsCount(repo domain.ProductRepository) *GetDiscountProductsCount {
	return &GetDiscountProductsCount{repo: repo}
}

func (uc *GetDiscountProductsCount) Execute(timeout time.Duration) (int, bool) {
	start := time.Now()
	
	for {
		currentCount := uc.repo.GetDiscountProductsCount()
		
		if time.Since(start) > timeout {
			return currentCount, false
		}
		
		time.Sleep(500 * time.Millisecond)
	}
}