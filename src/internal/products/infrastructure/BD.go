package infrastructure

import (
	"recu/src/internal/products/domain"
	"sync"
	"time"
)

type MemoryProductRepository struct {
	products       []domain.Product
	mu             sync.Mutex
	newProducts    int 
	discountCount  int 
	lastChecked    time.Time
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products:      make([]domain.Product, 0),
		newProducts:   0,
		discountCount: 0,
		lastChecked:   time.Now(),
	}
}

func (r *MemoryProductRepository) CreateProduct(product domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	r.products = append(r.products, product)
	r.newProducts++
	
	if product.Descuento {
		r.discountCount++
	}
	
	return nil
}

func (r *MemoryProductRepository) GetNewProductsCount() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	count := r.newProducts
	r.newProducts = 0 
	return count
}

func (r *MemoryProductRepository) GetDiscountProductsCount() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	return r.discountCount
}