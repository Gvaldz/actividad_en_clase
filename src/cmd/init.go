package cmd

import (
	productDeps  "recu/src/internal/products/infrastructure"
	"recu/src/server"
)

func Init() {

	productDependencies := productDeps.NewProductDependences()
	productRoutes := productDependencies.GetRoutes()

	server.Run(productRoutes)
}
