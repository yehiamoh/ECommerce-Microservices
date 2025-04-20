package handlers

import (
	"context"
	pb "product-service/gen/product"
	"product-service/pkg/services"
	"strconv"
)

type ProductHandler struct {
	productService *services.ProductService	
	pb.UnimplementedProductserviceServer
}
func NewProductHandler(svc *services.ProductService)*ProductHandler{
	return&ProductHandler{productService: svc}
}
func(h *ProductHandler)GetProduct(ctx context.Context,req *pb.GetProductRequest)(*pb.ProductResponse,error){
	ID,err:=strconv.Atoi(req.Id)
	if err!=nil{
		return nil,err
	}
	return h.productService.GetProductByIDService(ID)
}
func(h *ProductHandler)CreateProduct(ctx context.Context,req *pb.CreateProductRequest)(*pb.ProductResponse,error){
	return h.productService.CreateProductService(req.Name,req.Description,req.Price)
}
func (h *ProductHandler)DeleteProduct(ctx context.Context,req *pb.DeleteProductRequest)(*pb.ProductResponse,error){
	ID,err:=strconv.Atoi(req.Id)
	if err!=nil{
		return nil,err
	}
	return h.productService.DeleteProductService(ID)
}
func (h *ProductHandler)UpdateProduct(ctx context.Context,req *pb.UpdateProductRequest)(*pb.ProductResponse,error){
	ID,err:=strconv.Atoi(req.Id)
	if err!=nil{
		return nil,err
	}
	return h.productService.UpdateProductService(ID,req.Name,req.Description,req.Price)
}
func (h*ProductHandler)GetAllProducts(ctx context.Context,req *pb.Empty)(*pb.AllProductResponse,error){
	return h.productService.GetAllProductService()
}