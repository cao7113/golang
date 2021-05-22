package server

import (
	"context"
	"flag"
	"fmt"
	pingv1 "github.com/cao7113/hellogolang/proto/gosdk/proto/ping/v1"
	streamv1 "github.com/cao7113/hellogolang/proto/gosdk/proto/stream/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc/reflection"
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

var ConnAddress = flag.String("ConnAddress", ":50051", "rpc address")

func StartRPCServer() {
	flag.Parse()

	address := *ConnAddress
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatalf("failed to listen addr %s with error: %v", address, err)
	}
	defer lis.Close()

	var opts []grpc.ServerOption
	opts = setupMiddlewares(opts)
	s := grpc.NewServer(opts...)
	pb.RegisterHelloServiceServer(s, &HelloServer{})
	pingv1.RegisterPingServiceServer(s, &PingServer{})
	streamv1.RegisterStreamServiceServer(s, &StreamServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// todo graceful shutdown!!! todo
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
