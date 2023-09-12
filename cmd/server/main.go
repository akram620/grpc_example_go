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

func (ProfileGrpcServer) GetUserInfoById(ctx context.Context, req *grpc_pb.UserId) (*grpc_pb.UserInfo, error) {
	fmt.Println("user id: ", req.UserId)

	rand.Seed(time.Now().UnixNano())

	return &grpc_pb.UserInfo{
		Id:   req.GetUserId(),
		Name: "Name",
		Age:  int32(rand.Intn(21) + 10),
	}, nil
}

func (ProfileGrpcServer) GetUserContentById(ctx context.Context, req *grpc_pb.UserId) (*grpc_pb.Contents, error) {
	fmt.Println("user id: ", req.UserId)

	rand.Seed(time.Now().UnixNano())

	content := []*grpc_pb.Content{
		{Id: 1, ContentType: grpc_pb.ContentTypes_image, Url: "www."},
		{Id: 2, ContentType: grpc_pb.ContentTypes_image, Url: "www."},
		{Id: 3, ContentType: grpc_pb.ContentTypes_video, Url: "www."},
	}

	return &grpc_pb.Contents{
		UserId:   req.UserId,
		Contents: content,
	}, nil
}
