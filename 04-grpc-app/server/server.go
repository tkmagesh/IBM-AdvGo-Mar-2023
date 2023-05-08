package main

import (
	"fmt"

	"github.com/tkmagesh/ibm-advgo-mar-2023/04-grpc-app/proto"
)

type appService struct {
	proto.UnimplementedAppServiceServer
}

func (*asi appService) Add(ctx context.Context, req *proto.AddRequest) (*AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Println("[appService.add() processing %d and %d\n", x, y)

	result := x + y
	res := &proto.AddResponse{
		Result : result,
	}
	return res, nil
}

func main() {

}
