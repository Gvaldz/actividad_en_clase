package domain

type ProductRepository interface {
	CreateProduct(Product) error
	GetNewProductsCount() int          
	GetDiscountProductsCount() int      
}