package controllers

import (
	"net/http"
	"recu/src/internal/products/application"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type GetDiscountProductsCountController struct {
	useCase *application.GetDiscountProductsCount
}

func NewGetDiscountProductsCountController(useCase *application.GetDiscountProductsCount) *GetDiscountProductsCountController {
	return &GetDiscountProductsCountController{useCase: useCase}
}

func (c *GetDiscountProductsCountController) Handle(ctx *gin.Context) {
	timeoutStr := ctx.DefaultQuery("timeout", "5") 
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid timeout value"})
		return
	}

	count, updated := c.useCase.Execute(time.Duration(timeout) * time.Second)
	
	response := gin.H{
		"discount_products_count": count,
		"data_updated":           updated,
	}
	
	ctx.JSON(http.StatusOK, response)
}