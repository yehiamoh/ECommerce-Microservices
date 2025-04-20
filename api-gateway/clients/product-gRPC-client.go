package client

import (
	pb "api-gateway/gen/product"
	"context"
	"time"

	"google.golang.org/grpc"
)
func productClient() (pb.ProductserviceClient,error) {
	conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure(),grpc.WithBlock())
	if err!=nil{
		return nil,err
	}
	 client:=pb.NewProductserviceClient(conn)
	 return client,nil
}
func GetAllProducts(page,limit int32)(*pb.AllProductResponse,error){
	client,err:=productClient()
	if err!=nil{
		return nil,err
	}
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	res,err:=client.GetAllProducts(ctx,&pb.GetAllProductRequest{
		Page: page,
		Limit: limit,
	})
	return res,err

}