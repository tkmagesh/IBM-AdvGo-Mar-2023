package main

import (
	"context"
	"fmt"
	"io"
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
	// doRequestResponse(ctx, appServiceClient)
	// doServerStreaming(ctx, appServiceClient)
	doClientStreaming(ctx, appServiceClient)

}

func doRequestResponse(ctx context.Context, appServiceClient proto.AppServiceClient) {
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

func doServerStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := appServiceClient.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		primeRes, err := clientStream.Recv()
		if err == io.EOF {
			log.Println("All prime number received")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Prime No : %d\n", primeRes.GetPrimeNo())
	}
}

func doClientStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	nos := []int32{3, 1, 4, 2, 5, 8, 6, 9, 7}
	clientStream, err := appServiceClient.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		req := &proto.AverageRequest{
			No: no,
		}
		log.Printf("Sending no : %d\n", no)
		clientStream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Average Response, count = %d & average = %d\n", res.GetCount(), res.GetAverage())
}
