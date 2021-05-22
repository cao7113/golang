package server

import (
	"context"
	"fmt"
	"github.com/cao7113/hellogolang/proto/gosdk/proto/try/v1"
	"github.com/sirupsen/logrus"
	"time"
)

type TryServer struct {
	tryv1.UnimplementedTryServiceServer
}

//func (h TryServer) Hello(ctx context.Context, req *tryv1.HelloRequest) (*tryv1.HelloResponse, error) {
//	logrus.Infof("handling request: %+v", req)
//	switch req.Error {
//	case "error":
//		st := status.New(codes.FailedPrecondition, "failed to satisfy pre-conditions")
//		ds, err := st.WithDetails(
//			&tryv1.Error{
//				Code:    int64(123),
//				Message: "mock error code",
//			},
//		)
//		if err != nil {
//			return nil, st.Err()
//		}
//		return nil, ds.Err()
//	}
//	reply := &tryv1.HelloResponse{
//		Message: fmt.Sprintf("hi %s", req.Name),
//	}
//	return reply, nil
//}

func (h TryServer) Try(ctx context.Context, req *tryv1.TryRequest) (*tryv1.TryResponse, error) {
	msg := req.Name
	switch req.Gender.(type) {
	case *tryv1.TryRequest_Male:
		msg += " Male"
	case *tryv1.TryRequest_Female:
		msg += " Female"
	}
	msg += fmt.Sprintf(" Score: %d", req.Score)
	resp := &tryv1.TryResponse{Message: msg}
	return resp, nil
}

func (h TryServer) Timeout(ctx context.Context, req *tryv1.TimeoutRequest) (*tryv1.TimeoutResponse, error) {
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
	resp := &tryv1.TimeoutResponse{
		Msg: fmt.Sprintf("try-timeout reply after %d ms", time.Since(t0).Milliseconds()),
	}
	return resp, nil
}

func (h TryServer) Slow(ctx context.Context, req *tryv1.SlowRequest) (*tryv1.SlowResponse, error) {
	logrus.Infof("[server] requesting with %+v", req)
	time.Sleep(time.Duration(req.Seconds) * time.Second)
	resp := &tryv1.SlowResponse{
		Msg: fmt.Sprintf("slow reply after %d seconds", req.Seconds),
	}
	return resp, nil
}

func (h TryServer) Context(ctx context.Context, req *tryv1.ContextRequest) (*tryv1.ContextResponse, error) {
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
	rp := &tryv1.ContextResponse{
		Msg: fmt.Sprintf("response from server for from: %s", req.From),
	}
	return rp, nil
}

func (h TryServer) Fatal(ctx context.Context, req *tryv1.FatalRequest) (*tryv1.FatalResponse, error) {
	logrus.Fatalf("[server] handling fatal request: %+v", req)
	return nil, nil
}

func fibN(n int) int {
	if n <= 1 {
		return n
	}
	return fibN(n-1) + fibN(n-2)
}
