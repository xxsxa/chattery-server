package v1

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	v1 "github.com/xxsxa/chattery-server/pkg/api/v1"
	"log"
)

type chatService struct {
	msg chan string
}

func NewChatServiceServer() v1.ChatServiceServer {
	return &chatService{msg: make(chan string, 1000)}
}

func (s *chatService) Send(ctx context.Context, message *wrappers.StringValue) (*empty.Empty, error) {
	if message != nil {
		log.Printf("Sent request: message=%v", *message)
		s.msg <- message.Value
	} else {
		log.Print("Sent requested: message=<empty>")
	}
	return &empty.Empty{}, nil
}

func (s *chatService) Subscribe(e *empty.Empty, stream v1.ChatService_SubscribeServer) error {
	log.Print("Subscribe requested")
	for {
		m := <-s.msg
		n := v1.Message{Text: fmt.Sprintf("I have received from you: %s. Thank!", m)}
		if err := stream.Send(&n); err != nil {
			s.msg <- m
			log.Printf("Stream connection failed: %v", err)
			return nil
		}
	}
}
