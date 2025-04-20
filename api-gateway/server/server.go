package server

import (
	"api-gateway/routes"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() error {
	r := gin.Default()
	r.GET("/products",routes.GetProductsRoute)
	if err:=r.Run(":8080");err!=nil{
		return err
	}
	return nil
}