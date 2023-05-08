package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tkmagesh/ibm-advgo-mar-2023/04-grpc-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	appServiceClient := proto.NewAppServiceClient(clientConn)

	ctx := context.Background()
	addReq := &proto.AddRequest{
		X: 100,
		Y: 200,
	}

	res, err := appServiceClient.Add(ctx, addReq)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Result = %d\n", res.GetResult())

}
