package services

import (
	client "api-gateway/clients"
	pb "api-gateway/gen/product"
	"context"
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