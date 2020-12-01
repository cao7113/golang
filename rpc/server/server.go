package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/cao7113/hellogolang/config"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	//grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
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

	var opts []grpc.ServerOption
	opts = setupMiddlewares(opts)
	s := grpc.NewServer(opts...)
	hello.RegisterGreeterServer(s, &HelloServer{})

	logrus.Infof("running grpc server at %s", address)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

var panicFunc = func(ctx context.Context, p interface{}) (err error) {
	wrappedError := fmt.Errorf("panic occurs %v", p)

	//txn := newrelic.FromContext(ctx)
	//if txn != nil {
	//	txn.NoticeError(wrappedError)
	//}

	logrus.Error(wrappedError)

	return status.Error(codes.Internal, "Internal Error")
}

func setupMiddlewares(opts []grpc.ServerOption) []grpc.ServerOption {
	var streamInterceptors []grpc.StreamServerInterceptor
	var unaryInterceptors []grpc.UnaryServerInterceptor

	streamInterceptors = append(streamInterceptors,
		//grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(panicFunc)),
		grpc_ctxtags.StreamServerInterceptor(),
		grpctrace.StreamServerInterceptor(grpctrace.WithServiceName(config.Config.ServiceName)),
	)
	unaryInterceptors = append(unaryInterceptors,
		//grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(panicFunc)),
		grpc_ctxtags.UnaryServerInterceptor(),
		grpctrace.UnaryServerInterceptor(grpctrace.WithServiceName(config.Config.ServiceName)),
	)

	opts = append(opts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)))
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)))

	return opts
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
