package infrastructure

import (
	"recu/src/internal/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	createProductController           *controllers.CreateProductController
	getNewProductsCountController     *controllers.GetNewProductsCountController
	getDiscountProductsCountController *controllers.GetDiscountProductsCountController
}

func NewProductRoutes(
	createCtrl *controllers.CreateProductController,
	newCountCtrl *controllers.GetNewProductsCountController,
	discountCountCtrl *controllers.GetDiscountProductsCountController,
) *ProductRoutes {
	return &ProductRoutes{
		createProductController:           createCtrl,
		getNewProductsCountController:     newCountCtrl,
		getDiscountProductsCountController: discountCountCtrl,
	}
}

func (r *ProductRoutes) SetupRoutes(router *gin.Engine) {
	productGroup := router.Group("/products")
	{
		productGroup.POST("/addProducts", r.createProductController.Create)
		productGroup.GET("/isNewProductAdded", r.getNewProductsCountController.Handle)          
		productGroup.GET("/countProductsInDiscount", r.getDiscountProductsCountController.Handle)
	}
}