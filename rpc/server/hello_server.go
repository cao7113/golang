package server

import (
	"context"
	"fmt"
	hellov1 "github.com/cao7113/hellogolang/proto/gosdk/proto/hello/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type HelloServer struct {
	hellov1.UnimplementedHelloServiceServer
}

func (h HelloServer) TryTimeout(ctx context.Context, req *hellov1.TryTimeoutRequest) (*hellov1.TryTimeoutResponse, error) {
	logrus.Infof("[server] handling request with %+v", req)
	t0 := time.Now()
	i := 0
	for { // 检测超时
		i++
		if ctx.Err() == context.Canceled {
			logrus.Warnf("hit cancelled after %d loop", i)
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	resp := &hellov1.TryTimeoutResponse{
		Msg: fmt.Sprintf("try-timeout reply after %d ms", time.Since(t0).Milliseconds()),
	}
	return resp, nil
}

func (h HelloServer) Slow(ctx context.Context, req *hellov1.SlowRequest) (*hellov1.SlowResponse, error) {
	logrus.Infof("[server] requesting with %+v", req)
	time.Sleep(time.Duration(req.Seconds) * time.Second)
	resp := &hellov1.SlowResponse{
		Msg: fmt.Sprintf("slow reply after %d seconds", req.Seconds),
	}
	return resp, nil
}

func (h HelloServer) TryContext(ctx context.Context, req *hellov1.TryContextRequest) (*hellov1.TryContextResponse, error) {
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
	rp := &hellov1.TryContextResponse{
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

func (h HelloServer) Hello(ctx context.Context, req *hellov1.HelloRequest) (*hellov1.HelloResponse, error) {
	logrus.Infof("handling request: %+v", req)
	switch req.Error {
	case "error":
		st := status.New(codes.FailedPrecondition, "failed to satisfy pre-conditions")
		ds, err := st.WithDetails(
			&hellov1.Error{
				Code:    int64(123),
				Message: "mock error code",
			},
		)
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	reply := &hellov1.HelloResponse{
		Message: fmt.Sprintf("hi %s", req.Name),
	}
	return reply, nil
}
