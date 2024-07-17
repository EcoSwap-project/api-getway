package main

import (
	"log"

	api "api_getway/api"
	handler "api_getway/api/handlers"
	"api_getway/config"
	auth "api_getway/genproto/authentication_service"
	reser "api_getway/genproto/item_service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mux := gin.Default()

	authServerConn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to authentication service: ", err)
	}
	defer authServerConn.Close()

	itemServerConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to reservation service: ", err)
	}
	defer itemServerConn.Close()

	authClient := auth.NewEcoServiceClient(authServerConn)
	itemClient := reser.NewEcoExchangeServiceClient(itemServerConn)

	handlers := handler.NewHandler(authClient, itemClient)
	log.Println("Starting API Gateway...")

	server := api.NewServer(handlers)

	server.InitRoutes(mux)

	mux.Run(":" + config.Load().URL_PORT)
}
