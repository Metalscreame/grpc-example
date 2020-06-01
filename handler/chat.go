package handler

import (
	"log"

	"github.com/Metalscreame/gRPC_example/pb"
	"golang.org/x/net/context"
)

type Chat struct {}

func (s *Chat) SayHello(_ context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message from client: %v", in)
	return &pb.Message{Body: "Hello From the Server:ChatHandler!"}, nil
}