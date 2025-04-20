package main

import (
	"api-gateway/server"
	"log"
)

func main() {
	if err:=server.SetUpRoutes();err!=nil{
		log.Fatal("Error in Running server : ",err)
	}
}