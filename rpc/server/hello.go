package server

import (
	"context"
	"fmt"
	pb "github.com/cao7113/hellogolang/proto/gosdk/proto/hello/v1"
	"github.com/sirupsen/logrus"
)

type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s HelloServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	logrus.Infof("[server] handling hello-request with %+v", req)
	resp := &pb.HelloResponse{
		Message: fmt.Sprintf("Welcome %s", req.From),
	}
	return resp, nil
}
