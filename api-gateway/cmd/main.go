package main

import (
	client "api-gateway/pkg/clients"
	"api-gateway/pkg/routes"
	"api-gateway/pkg/server"
	"api-gateway/pkg/services"
	"log"
)

func main() {
	productClient, err := client.NewProductClient()
    if err != nil {
        log.Fatal("Failed to connect to ProductService:", err)
    }

    // Initialize services
    productService := services.NewProductService(productClient)

    // Initialize routes
    productRoutes := routes.NewProductRoutes(productService)

    // Register routes
    server := server.NewServer()
    server.RegisterRoutes(productRoutes)

    // Start server
    if err := server.Run(); err != nil {
        log.Fatal("Error in running server:", err)
    }
}