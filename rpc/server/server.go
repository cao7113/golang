package server

import (
	"context"
	"flag"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	//grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	pb "github.com/cao7113/hellogolang/proto/gosdk/proto/hello/v1"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"net"
)

var ConnAddress = flag.String("ConnAddress", "localhost:50051", "rpc address")

//var (
//	_ pb.GreeterServer = &HelloServer{}
//)

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
	svr := &HelloServer{}
	pb.RegisterHelloServiceServer(s, svr)

	logrus.Infof("running grpc server at %s", address)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func setupMiddlewares(opts []grpc.ServerOption) []grpc.ServerOption {
	streamInterceptors := []grpc.StreamServerInterceptor{
		grpc_ctxtags.StreamServerInterceptor(),
		grpctrace.StreamServerInterceptor(),
	}
	opts = append(opts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)))

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		//grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(panicFunc)),
		grpc_ctxtags.UnaryServerInterceptor(),
		grpctrace.UnaryServerInterceptor(),
	}
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)))

	return opts
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
