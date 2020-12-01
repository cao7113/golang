package server

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/cao7113/hellogolang/rpc/hello"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ConnAddress = flag.String("ConnAddress", "localhost:50051", "rpc address")

func StartRPCServer() {
	flag.Parse()

	address := *ConnAddress
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &HelloServer{})

	logrus.Infof("running grpc server at %s", address)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

type HelloServer struct{}

func (h HelloServer) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloReply, error) {
	logrus.Infof("handling request: %+v", req)
	switch req.Error {
	case "error":
		st := status.New(codes.FailedPrecondition, "failed to satisfy pre-conditions")
		ds, err := st.WithDetails(
			&hello.Error{
				Code:    int64(123),
				Message: "mock error code",
			},
		)
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	reply := &hello.HelloReply{
		Message: fmt.Sprintf("hi %s", req.Name),
	}
	return reply, nil
}
