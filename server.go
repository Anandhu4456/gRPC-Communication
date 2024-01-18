package main

import (
	"go-grpc/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("cant listen on port 50005: %v", err)
	}
	s := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatSeviceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server grpc server on port :50005 %v", err)
	}
}
