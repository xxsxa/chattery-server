package grpc

import (
	"context"
	v1 "github.com/xxsxa/chattery-server/pkg/api/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunServer(ctx context.Context,srv v1.ChatServiceServer,port string)error  {
	listen,err := net.Listen("tcp",":"+port)
	if err != nil{
		return err
	}
	server := grpc.NewServer()
	v1.RegisterChatServiceServer(server,srv)
	log.Println("starting server...")
	return server.Serve(listen)
}