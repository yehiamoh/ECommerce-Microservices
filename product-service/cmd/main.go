package main

import (
	"log"
	"product-service/pkg/server"
)

func main() {
	if err := server.Start();err!=nil{
		log.Fatalf("Error in Starting the server :%v",err)
	}
}