package services

import (
	pb "api-gateway/gen/product"
	client "api-gateway/pkg/clients"
	"context"
	"fmt"
	"strconv"
	"time"
)

func GetAllProductService(page, limit int32) (*pb.AllProductResponse, error) {
	client, err := client.ProductClient()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := client.GetAllProducts(ctx, &pb.GetAllProductRequest{
		Page:  page,
		Limit: limit,
	})
	return res, err

}
func GetProductByIdService(id int) (*pb.ProductResponse, error) {
	client, err := client.ProductClient()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	strId := strconv.Itoa(id)
	res, err := client.GetProduct(ctx, &pb.GetProductRequest{
		Id: strId,
	})
	return res, err
}
func CreateProductService(name string,description string,price float32)(*pb.ProductResponse,error){
	client,err:=client.ProductClient()
	if err!=nil{
		return nil,fmt.Errorf("error: in creating client (CreateProductService) :%v",err)
	}
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	res,err:=client.CreateProduct(ctx,&pb.CreateProductRequest{
			Name: name,
			Description: description,
			Price: price,})
	if err!=nil{
		return nil,fmt.Errorf("error: in creating product client Function :%v",err)
	}
	return res,nil
}
func DeleteProductService(id int)(*pb.ProductResponse, error){
	client, err := client.ProductClient()
	if err != nil {
		return nil,fmt.Errorf("error: in creating client (DeleteProductService) :%v",err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	strId := strconv.Itoa(id)
	res, err := client.DeleteProduct(ctx, &pb.DeleteProductRequest{
		Id: strId,
	})
	return res, err
}