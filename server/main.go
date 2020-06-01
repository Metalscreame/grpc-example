package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Metalscreame/gRPC_example/handler"
	"github.com/Metalscreame/gRPC_example/models/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	chatHandler := handler.Chat{}
	// registering specific handlers for this server
	chat.RegisterChatServiceServer(grpcServer, &chatHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
