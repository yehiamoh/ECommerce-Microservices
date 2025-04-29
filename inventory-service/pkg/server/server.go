package server

import (
	"fmt"
	pb "inventory-service/gen/inventory"
	"inventory-service/pkg/handlers"
	"inventory-service/pkg/repository"
	"inventory-service/pkg/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() error {
	db, err := repository.Open()
	if err!=nil{
		return fmt.Errorf("error in starting database:%v",err)
	}

	lis,err:=net.Listen("tcp","localhost:50053")
	if err!=nil{
		return fmt.Errorf("error in listening on port 50053:%v",err)
	}

	inventoryRepo:=repository.NewInventoryRepository(db)
	inventoryService:=services.NewInvenotryservice(inventoryRepo)
	inventoryHandlers:=handlers.NewInventoryHandlers(inventoryService)

	grpcServer:= grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterInventoryServiceServer(grpcServer,inventoryHandlers)
	
	log.Println("server started")
	
	grpcServer.Serve(lis)

	return nil
}