package rpc

import (
	"context"
	"fmt"
	pb "github.com/cao7113/hellogolang/rpc/protos"
	"github.com/cao7113/hellogolang/rpc/server"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

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
	for i := 0; i < 1; i++ {
		for j := 0; j < 10; j++ {
			time.Sleep(1 * time.Millisecond)
			go func(i, j int) {
				ctx := context.Background()
				c := pb.NewGreeterClient(conn)
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &pb.TcRequest{
					From: from,
				}
				r, err := c.TryContext(ctx, req)
				logrus.Infof("client request with from %s", from)
				s.Nil(err)
				logrus.Infof("response: %s", r.Msg)
			}(i, j)
		}
	}

	time.Sleep(10 * time.Hour)
	logrus.Infof("over")
}

func (s *ClientTestSuite) TestWithoutTimeout() {
	conn, err := grpc.Dial(*server.ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			logrus.Infof("failed to close connection: %s", e)
		}
	}()

	// no timeout
	for i := 0; i < 100; i++ {
		for j := 0; j < 500; j++ {
			time.Sleep(1 * time.Millisecond)
			go func(i, j int) {
				ctx := context.Background()
				c := pb.NewGreeterClient(conn)
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &pb.TcRequest{
					From: from,
				}
				r, err := c.TryContext(ctx, req)
				logrus.Infof("client request with from %s", from)
				s.Nil(err)
				logrus.Infof("response: %s", r.Msg)
			}(i, j)
		}
	}

	time.Sleep(30 * time.Hour)
	logrus.Infof("over")
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

	c := pb.NewGreeterClient(conn)
	req := &pb.HelloRequest{
		Name:  "try",
		Error: "error",
	}
	r, err := c.SayHello(ctx, req)
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
	logrus.Errorf("protos error: %+v", err)
}

type ClientTestSuite struct {
	suite.Suite
}

// The SetupSuite method will be run before any tests are run.
func (s *ClientTestSuite) SetupSuite() {
	//runtime.GOMAXPROCS(2)
	go server.StartRPCServer()
	time.Sleep(1 * time.Second)
}

// The TearDownSuite method will be run after all tests have been run.
func (s *ClientTestSuite) TearDownSuite() {
}

func TestClientRequests(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
