package main

import (
	"context"
	"errors"
	"fmt"
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
