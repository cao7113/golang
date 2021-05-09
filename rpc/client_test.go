package rpc

import (
	"context"
	"fmt"
	pb "github.com/cao7113/hellogolang/proto/gosdk/proto/hello/v1"
	"github.com/cao7113/hellogolang/rpc/server"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
	"testing"
	"time"
)

// call with timeout context

func (s *ClientTestSuite) TestWithCallTimeout() {
	ct := context.Background()
	bct, cancel := context.WithTimeout(ct, 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(bct, *server.ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			logrus.Infof("failed to close connection: %s", e)
		}
	}()

	// without context timeout
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			//time.Sleep(1 * time.Millisecond)
			go func(i, j int) {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
				defer cancel()
				c := pb.NewHelloServiceClient(conn)
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &pb.TryContextRequest{
					From: from,
				}
				logrus.Infof("client request with from %s", from)
				r, err := c.TryContext(ctx, req)
				if err != nil {
					st := status.Convert(err)
					logrus.Errorf("from: %s error: code: %d message: %s", from, st.Code(), st.Message())
				}
				if r != nil {
					logrus.Infof("client from: %s got msg: %s", from, r.Msg)
				}
			}(i, j)
		}
	}

	time.Sleep(10 * time.Hour)
	logrus.Infof("over")
}

func (s *ClientTestSuite) TestWithTimeout() {
	ct := context.Background()
	bct, cancel := context.WithTimeout(ct, 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(bct, *server.ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			logrus.Infof("failed to close connection: %s", e)
		}
	}()

	// without context timeout
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			//time.Sleep(1 * time.Millisecond)
			go func(i, j int) {
				ctx := context.Background()
				c := pb.NewHelloServiceClient(conn)
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &pb.TryContextRequest{
					From: from,
				}
				logrus.Infof("client request with from %s", from)
				r, err := c.TryContext(ctx, req)
				if err != nil {
					st := status.Convert(err)
					logrus.Errorf("from: %s error: code: %d message: %s", from, st.Code(), st.Message())
				}
				if r != nil {
					logrus.Infof("client from: %s got msg: %s", from, r.Msg)
				}
			}(i, j)
		}
	}

	time.Sleep(10 * time.Hour)
	logrus.Infof("over")
}

func (s *ClientTestSuite) TestWithTimeout1() {
	ct := context.Background()
	bct, cancel := context.WithTimeout(ct, 2*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(bct, *server.ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			logrus.Infof("failed to close connection: %s", e)
		}
	}()

	time.Sleep(3 * time.Second) // let dial context timeout

	//ctx := context.Background()
	ctx := bct
	c := pb.NewHelloServiceClient(conn)
	for i := 0; i < 1; i++ {
		for j := 0; j < 1; j++ {
			//time.Sleep(1 * time.Millisecond)
			go func(i, j int) {
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &pb.TryContextRequest{
					From: from,
				}
				logrus.Infof("client request with from %s", from)
				r, err := c.TryContext(ctx, req)
				if err != nil {
					st := status.Convert(err)
					logrus.Errorf("from: %s error: code: %d message: %s", from, st.Code(), st.Message())
				}
				if r != nil {
					logrus.Infof("client from: %s got msg: %s", from, r.Msg)
				}
			}(i, j)
		}
	}

	time.Sleep(10 * time.Hour)
	logrus.Infof("over")
}

func (s *ClientTestSuite) TestAllWithoutTimeout() {
	conn, err := grpc.Dial(*server.ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			logrus.Infof("failed to close connection: %s", e)
		}
	}()

	for i := 0; i < 200; i++ {
		for j := 0; j < 3; j++ {
			//time.Sleep(1 * time.Millisecond)
			go func(i, j int) {
				ctx := context.Background()
				c := pb.NewHelloServiceClient(conn)
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &pb.TryContextRequest{
					From: from,
				}
				logrus.Infof("client request with from %s", from)
				r, err := c.TryContext(ctx, req)
				if err != nil {
					st := status.Convert(err)
					logrus.Errorf("from: %s error: code: %d message: %s", from, st.Code(), st.Message())
				}
				if r != nil {
					logrus.Infof("client from: %s got msg: %s", from, r.Msg)
				}
			}(i, j)
		}
	}

	time.Sleep(10 * time.Hour)
	logrus.Infof("over")
}

func (s *ClientTestSuite) TestStatusCode() {
	reason := "testing"
	err := status.Errorf(codes.FailedPrecondition, "failed precondition %s", reason)

	err1, ok := status.FromError(err)
	s.True(ok)
	s.Equal(codes.FailedPrecondition, err1.Code())
	// "msg: failed precondition testing, code: FailedPrecondition, str: FailedPrecondition"
	logrus.Infof("msg: %s, code: %v, str: %s", err1.Message(), err1.Code(), err1.Code().String())
}

func (s *ClientTestSuite) TestDetailError() {
	conn, err := grpc.Dial(*server.ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			logrus.Infof("failed to close connection: %s", e)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c := pb.NewHelloServiceClient(conn)
	req := &pb.HelloRequest{
		Name:  "try",
		Error: "error",
	}
	r, err := c.Hello(ctx, req)
	s.Nil(r)
	s.NotNil(err)
	se := status.Convert(err)
	for _, d := range se.Details() {
		switch info := d.(type) {
		case *pb.Error:
			logrus.Infof("hit mock error: %+v", info)
		default:
			logrus.Fatalf("Unexpected type: %s", info)
		}
	}
	logrus.Errorf("proto error: %+v", err)
}

type ClientTestSuite struct {
	suite.Suite
}

// The SetupSuite method will be run before any tests are run.
func (s *ClientTestSuite) SetupSuite() {
	runtime.GOMAXPROCS(2)
	go server.StartRPCServer()
	time.Sleep(1 * time.Second)
}

// The TearDownSuite method will be run after all tests have been run.
func (s *ClientTestSuite) TearDownSuite() {
}

func TestClientRequests(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
