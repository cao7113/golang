package server

import (
	"context"
	"fmt"
	"github.com/cao7113/hellogolang/rpc/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HelloServer struct{}

func (h HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	logrus.Infof("handling request: %+v", req)
	switch req.Error {
	case "error":
		st := status.New(codes.FailedPrecondition, "failed to satisfy pre-conditions")
		ds, err := st.WithDetails(
			&pb.Error{
				Code:    int64(123),
				Message: "mock error code",
			},
		)
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	reply := &pb.HelloReply{
		Message: fmt.Sprintf("hi %s", req.Name),
	}
	return reply, nil
}
