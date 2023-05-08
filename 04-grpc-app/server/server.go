package main

import (
	"context"
	"fmt"
	"log"
	"net"

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

	result := x + y
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
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
