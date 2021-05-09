package server

import (
	"context"
	"fmt"
	pb "github.com/cao7113/hellogolang/rpc/protos/gen/protos"
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
	if ok {
		logrus.Infof("request %+v with deadline: %s", req, dt)
	}
	i := 0
	for {
		n := fibN(40) // 102334155
		time.Sleep(10 * time.Millisecond)

		//n := fibN(1)
		//time.Sleep(2 * time.Second)

		logrus.Infof("%s requesting from: %s count: %+v fibN: %d", time.Now().Format(time.RFC3339), req.From, i, n)
		i++
	}
	rp := &pb.TcReply{
		Msg: fmt.Sprintf("response from server for from: %s", req.From),
	}
	return rp, nil
}

func fibN(n int) int {
	if n <= 1 {
		return n
	}
	return fibN(n-1) + fibN(n-2)
}

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
