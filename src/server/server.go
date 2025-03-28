package server

import (
	productsRoutes "recu/src/internal/products/infrastructure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(
	productsRoutes *productsRoutes.ProductRoutes,
) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	
	}))

	productsRoutes.SetupRoutes(r)

	r.Run(":8080")
}
