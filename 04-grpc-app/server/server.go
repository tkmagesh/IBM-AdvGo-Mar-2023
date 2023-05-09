package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/tkmagesh/ibm-advgo-mar-2023/04-grpc-app/proto"
	"google.golang.org/grpc"
)

type appService struct {
	proto.UnimplementedAppServiceServer
}

func (asi *appService) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("[appService.add()] processing %d and %d\n", x, y)

	time.Sleep(5 * time.Second) // forcing the timeout signal to be triggerd from the client
	select {
	case <-ctx.Done():
		log.Println("timeout occurred")
		return nil, errors.New("timeout occurred")
	default:
		result := x + y
		res := &proto.AddResponse{
			Result: result,
		}
		log.Println("sending response")
		return res, nil
	}

}

func (asi *appService) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	log.Printf("Generating primes, start = %d and end = %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			log.Printf("Sending prime no : %d\n", no)
			serverStream.Send(res)
			time.Sleep(500 * time.Millisecond)
		}
	}
	log.Println("All the prime numbers are generated")
	return nil

}

func (asi *appService) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var count, sum int32
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			log.Println("Received all the values")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Received value : %d\n", req.GetNo())
		count++
		sum += req.GetNo()
	}

	avg := sum / count
	res := &proto.AverageResponse{
		Count:   count,
		Average: avg,
	}
	log.Println("Sending the average result")
	if err := serverStream.SendAndClose(res); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	asi := &appService{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
