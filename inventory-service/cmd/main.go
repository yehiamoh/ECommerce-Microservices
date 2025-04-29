package main

import (
	"inventory-service/pkg/server"
	"log"
)

func main() {
	if err:=server.Start();err!=nil{
		log.Fatalf("error in starting the server: %v",err)
	}
}