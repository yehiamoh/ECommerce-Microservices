package server

import (
	"api-gateway/pkg/routes"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() error {
	r := gin.Default()
	productRoutes(r)
	if err:=r.Run(":8080");err!=nil{
		return err
	}
	return nil
}
func productRoutes(router *gin.Engine){
	router.GET("/products",routes.GetProductsRoute)
	router.GET("/products/:id",routes.GetProductByIdRoute)
} 