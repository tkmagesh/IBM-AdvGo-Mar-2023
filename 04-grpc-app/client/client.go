package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tkmagesh/ibm-advgo-mar-2023/04-grpc-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	appServiceClient := proto.NewAppServiceClient(clientConn)

	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	addReq := &proto.AddRequest{
		X: 100,
		Y: 200,
	}

	res, err := appServiceClient.Add(timeoutCtx, addReq)
	if err != nil {
		if code := status.Code(err); code == codes.DeadlineExceeded {
			fmt.Println("timeout occurred")
			return
		}
		log.Fatalln(err)
	}
	fmt.Printf("Result = %d\n", res.GetResult())

}
