package server

import (
	"fmt"
	"log"
	"net"
	"user-service/pkg/handlers"
	"user-service/pkg/repository"
	"user-service/pkg/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "user-service/gen/user"
)

func Start()error{
	db,err:=repository.Open()
	if err!=nil{
		return fmt.Errorf("error in starting database:%v",err)
	}

	defer db.Close()
	repo:=repository.NewUserRepository(db)
	svc:=service.NewUserService(repo)
	handlers:=handlers.NewUserHandlers(svc)

	lis,err:=net.Listen("tcp","localhost:50052")
	if err!=nil{
		return fmt.Errorf("error in listening on port 50052:%v",err)
	}

	grpcServer:=grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterUserServiceServer(grpcServer,handlers)

	log.Println("server started")
	
	grpcServer.Serve(lis)

	return nil

}