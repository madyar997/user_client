package server

import (
	"context"
	"log"
	"user-client/protobuf"
)

type MyGrpcServer struct {
	protobuf.UnimplementedUserServer
}

func (m *MyGrpcServer) GetUserByID(ctx context.Context, req *protobuf.UserRequest) (*protobuf.UserResponse, error) {
	log.Println("we are inside get user by id method")

	return &protobuf.UserResponse{
		Id:    1,
		Name:  "Madyar",
		Email: "madiar.997@gmail.com",
		Age:   25,
	}, nil
}
