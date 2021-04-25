package server

import (
	"context"
	"fmt"
	pb "github.com/cao7113/hellogolang/rpc/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type HelloServer struct {
	pb.UnimplementedGreeterServer
}

func (h HelloServer) TryContext(ctx context.Context, req *pb.TcRequest) (*pb.TcReply, error) {
	dt, ok := ctx.Deadline()
	logrus.Infof("request %+v with deadline: %s ok: %+v in context: %+v", req, dt, ok, ctx)
	i := 0
	for {
		logrus.Infof("count: %+v", i)
		time.Sleep(2 * time.Second)
	}
	rp := &pb.TcReply{
		Msg: "response from server",
	}
	return rp, nil
}

//func (h HelloServer) mustEmbedUnimplementedGreeterServer() {
//	panic("implement me")
//}

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
