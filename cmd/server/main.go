package main

import (
	"context"
	"fmt"
	"github.com/akram620/grpc_example_go/pkg/grpc_pb"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"time"
)

func main() {
	s := grpc.NewServer()
	srv := &ProfileGrpcServer{}

	grpc_pb.RegisterProfileServer(s, srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}

	if err = s.Serve(l); err != nil {
		fmt.Println(err)
	}
}

type ProfileGrpcServer struct {
	grpc_pb.UnimplementedProfileServer
}

func (ProfileGrpcServer) GetUserInfo(ctx context.Context, req *grpc_pb.UserRequest) (*grpc_pb.UserResponse, error) {
	fmt.Println("Request user id: ", req.UserId)

	rand.Seed(time.Now().UnixNano())

	return &grpc_pb.UserResponse{
		Id:       req.GetUserId(),
		Name:     "Name",
		Age:      int32(rand.Intn(21) + 10),
		UserType: grpc_pb.UserTypes_teacher,
	}, nil
}
