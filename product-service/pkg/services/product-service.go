package services

import (
	pb "product-service/gen/product"
	"product-service/pkg/repository"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}
func NewProductService(repo *repository.ProductRepository)*ProductService{
	return &ProductService{
		ProductRepo: repo,
	}
}
func(s *ProductService)GetProductByIDService(id int)(*pb.ProductResponse,error){
	return s.ProductRepo.GetProductByID(id)
}
func(s *ProductService)CreateProductService(name string,description string,price float32)(*pb.ProductResponse,error){
	return s.ProductRepo.CreateProduct(name,description,price)
}
func(s*ProductService)UpdateProductService(id int, name string,description string,price float32)(*pb.ProductResponse,error){
	return s.ProductRepo.UpdateProduct(id,&name,&description,&price)
}
func(s *ProductService)DeleteProductService(id int)(*pb.ProductResponse,error){
	return s.ProductRepo.DeleteProduct(id)
}
func (s*ProductService)GetAllProductService()(*pb.AllProductResponse,error){
	return s.ProductRepo.GetAllProducts()
}