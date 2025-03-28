package controllers

import (
	"net/http"
	"recu/src/internal/products/application"

	"github.com/gin-gonic/gin"
)

type GetNewProductsCountController struct {
	useCase *application.GetNewProductsCount
}

func NewGetNewProductsCountController(useCase *application.GetNewProductsCount) *GetNewProductsCountController {
	return &GetNewProductsCountController{useCase: useCase}
}

func (c *GetNewProductsCountController) Handle(ctx *gin.Context) {
	count := c.useCase.Execute()
	ctx.JSON(http.StatusOK, gin.H{"new_products_count": count})
}