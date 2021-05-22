package server

import (
	"context"
	"fmt"
	pb "github.com/cao7113/hellogolang/proto/gosdk/proto/hello/v1"
)

type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (h HelloServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := &pb.HelloResponse{
		Message: fmt.Sprintf("Welcome %s", req.From),
	}
	return resp, nil
}
