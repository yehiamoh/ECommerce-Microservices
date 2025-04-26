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
        log.Fatal("Failed to connect to ProductCleint:", err)
    }
	userClient, err := client.NewUserClient()
    if err != nil {
        log.Fatal("Failed to connect to UserCleint:", err)
    }

    // Initialize services
    productService := services.NewProductService(productClient)
    userService :=services.NewUserService(userClient)
    // Initialize routes
    productRoutes := routes.NewProductRoutes(productService)
    userRoutes := routes.NewUserRoutes(userService)

    // Register routes
    server := server.NewServer()
    server.RegisterProductRoutes(productRoutes)
    server.RegitserUserRoutes(userRoutes)

    // Start server
    if err := server.Run(); err != nil {
        log.Fatal("Error in running server:", err)
    }
}