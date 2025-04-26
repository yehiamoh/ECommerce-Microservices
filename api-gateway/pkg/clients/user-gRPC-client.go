package client

import (
	pb "api-gateway/gen/user"

	"google.golang.org/grpc"
)
func NewUserClient()(pb.UserServiceClient,error){
	conn,err:=grpc.Dial("localhost:50052",grpc.WithInsecure(),grpc.WithBlock())
	if err!=nil{
		return nil,err
	}
	client:=pb.NewUserServiceClient(conn)
	return client,nil
}