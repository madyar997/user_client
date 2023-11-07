package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user-client/protobuf"
	"user-client/server"
)

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myUserService := &server.MyGrpcServer{}
	protobuf.RegisterUserServer(s, myUserService)

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
