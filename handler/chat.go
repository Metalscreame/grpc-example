package handler

import (
	"log"

	"github.com/Metalscreame/gRPC_example/models/chat"
	"golang.org/x/net/context"
)

type Chat struct {}

func (s *Chat) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &chat.Message{Body: "Hello From the Server:ChatHandler!"}, nil
}