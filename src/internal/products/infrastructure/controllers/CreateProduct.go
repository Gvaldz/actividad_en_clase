package controllers

import (
	"recu/src/internal/products/application"
	"recu/src/internal/products/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	CreateProduct *application.CreateProduct
}

func NewCreateProductController(create *application.CreateProduct) *CreateProductController {
	return &CreateProductController{
		CreateProduct: create,
	}
}

func (h *CreateProductController) Create(c *gin.Context) {
	var productRequest struct {
		Codigo      int32   `json:"codigo"`
		Nombre  	string  `json:"nombre"`
		Precio		float32 `json:"precio"`
		Descuento	bool    `json:"descuento"`

	}

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := domain.Product{
		Codigo:      productRequest.Codigo,
		Nombre:      productRequest.Nombre,
		Precio:      productRequest.Precio,
		Descuento:   productRequest.Descuento,
	}
	
	err := h.CreateProduct.Execute(product)
		if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
    c.JSON(http.StatusCreated, gin.H{"message": "Producto creado correctamente", "producto": product})
}