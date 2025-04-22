package client

import (
	pb "api-gateway/gen/product"

	"google.golang.org/grpc"
)
func NewProductClient() (pb.ProductserviceClient,error) {
	conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure(),grpc.WithBlock())
	if err!=nil{
		return nil,err
	}
	 client:=pb.NewProductserviceClient(conn)
	 return client,nil
}
