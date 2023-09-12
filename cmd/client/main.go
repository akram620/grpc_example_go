package main

import (
	"context"
	"fmt"
	"github.com/akram620/grpc_example_go/pkg/grpc_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer con.Close()

	client := grpc_pb.NewProfileClient(con)

	userInfo, err := client.GetUserInfoById(context.Background(), &grpc_pb.UserId{UserId: 1})
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%+v\n", userInfo)

	time.Sleep(time.Second)

	userContents, err := client.GetUserContentsById(context.Background(), &grpc_pb.UserId{UserId: 1})
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%+v\n", userContents)

}
