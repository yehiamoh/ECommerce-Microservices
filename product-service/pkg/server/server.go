package server

import (
	"log"
	"net"
	"product-service/pkg/handlers"
	"product-service/pkg/repository"
	"product-service/pkg/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "product-service/gen/product"
)

func Start() error {
	db, err := repository.Open()
	if err!=nil{
		log.Fatal("Opening Database",err)
		return err
	}

	repo:=repository.NewProductRepository(db)
	service:=services.NewProductService(repo)
	handler:=handlers.NewProductHandler(service)

	lis,err:=net.Listen("tcp","localhost:50051")
	if err!=nil{
		log.Fatal("Error in listening on port 50051",err)
		return err
	}

	grpcServer:=grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterProductserviceServer(grpcServer,handler)

	log.Println("server started")
	
	grpcServer.Serve(lis)

	return nil
}