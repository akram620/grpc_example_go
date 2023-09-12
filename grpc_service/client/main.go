package main

import (
	"context"
	"fmt"
	"github.com/akram620/grpc_example_go/pkg/grpc_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const address = "localhost:8081"

func main() {
	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer con.Close()

	client := grpc_api.NewProfileClient(con)

	for i := 1; i <= 10; i++ {
		response, err := client.GetUserInfo(context.Background(), &grpc_api.UserRequest{UserId: int64(i)})
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Printf("%+v\n", response)
		time.Sleep(time.Second)
	}

}