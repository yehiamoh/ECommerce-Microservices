package services

import (
	pb "api-gateway/gen/product"
	"context"
	"fmt"
	"strconv"
	"time"
)
type ProductService struct{
	client pb.ProductserviceClient
}
func NewProductService(client pb.ProductserviceClient)*ProductService{
	return &ProductService{
		client: client,
	}
}
func(s *ProductService)GetAllProductService(page, limit int32) (*pb.AllProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := s.client.GetAllProducts(ctx, &pb.GetAllProductRequest{
		Page:  page,
		Limit: limit,
	})
	return res, err

}
func(s *ProductService) GetProductByIdService(id int) (*pb.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	strId := strconv.Itoa(id)
	res, err := s.client.GetProduct(ctx, &pb.GetProductRequest{
		Id: strId,
	})
	return res, err
}
func (s *ProductService) CreateProductService(name string,description string,price float32)(*pb.ProductResponse,error){
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	res,err:=s.client.CreateProduct(ctx,&pb.CreateProductRequest{
			Name: name,
			Description: description,
			Price: price,})
	if err!=nil{
		return nil,fmt.Errorf("error: in creating product client Function :%v",err)
	}
	return res,nil
}
func(s *ProductService) DeleteProductService(id int)(*pb.ProductResponse, error){

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	strId := strconv.Itoa(id)
	res, err := s.client.DeleteProduct(ctx, &pb.DeleteProductRequest{
		Id: strId,
	})
	return res, err
}