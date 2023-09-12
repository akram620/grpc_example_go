package main

import (
	"context"
	"fmt"
	"github.com/akram620/grpc_example_go/pkg/grpc_api"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"time"
)

func main() {
	s := grpc.NewServer()
	srv := &ProfileGrpcServer{}

	grpc_api.RegisterProfileServer(s, srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}

	if err = s.Serve(l); err != nil {
		fmt.Println(err)
	}
}

type ProfileGrpcServer struct {
	grpc_api.UnimplementedProfileServer
}

func (ProfileGrpcServer) GetUserInfo(ctx context.Context, req *grpc_api.UserRequest) (*grpc_api.UserResponse, error) {
	fmt.Println("Request user id: ", req.UserId)

	rand.Seed(time.Now().UnixNano())

	return &grpc_api.UserResponse{
		Id:       req.GetUserId(),
		Name:     "Name",
		Age:      int32(rand.Intn(21) + 10),
		UserType: grpc_api.UserTypes_teacher,
	}, nil
}

func (ProfileGrpcServer) GetUserInfoV2(ctx context.Context, req *grpc_api.UserRequest) (*grpc_api.UserResponseV2, error) {
	return &grpc_api.UserResponseV2{
		Id:      req.GetUserId(),
		Name:    "Name",
		SurName: "A",
	}, nil
}
