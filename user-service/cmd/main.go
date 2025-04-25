package main

import (
	"log"
	"user-service/pkg/server"
)

func main() {
	if err:=server.Start();err!=nil{
		log.Fatalf("Error in starting the server:%v",err)
	}
}