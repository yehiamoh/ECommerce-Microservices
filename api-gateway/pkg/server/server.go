package server

import (
	"api-gateway/pkg/routes"

	"github.com/gin-gonic/gin"
)

type Server struct{
	Router *gin.Engine
}

func NewServer()*Server{
	return &Server{
		Router: gin.Default(),
	}
}

func (s *Server) RegisterProductRoutes(productRoutes *routes.ProductRoutes) {
    s.Router.GET("/products", productRoutes.GetProductsRoute)
	s.Router.POST("/products", productRoutes.CreateProductRoute)
    s.Router.GET("/products/:id", productRoutes.GetProductByIdRoute)
    s.Router.DELETE("/products/:id", productRoutes.DeleteProductRoute)
}
func (s *Server) RegitserUserRoutes(userRoutes *routes.UserRoutes) {
    s.Router.POST("/users", userRoutes.RegisterRoute)
    s.Router.POST("/users/login", userRoutes.LoginRoute)
	s.Router.GET("/users/:id",userRoutes.GetUserByID)
	
}

func (s *Server) Run() error {
    return s.Router.Run(":8080")
}