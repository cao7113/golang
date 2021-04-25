package rpc

import (
	"context"
	pb "github.com/cao7113/hellogolang/rpc/protos"
	"github.com/cao7113/hellogolang/rpc/server"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func (s *ClientTestSuite) TestTryCtx() {
	//ct := context.Background()
	//bct := context.WithTimeout(ct, 3 * time.Second)
	conn, err := grpc.Dial(*server.ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			logrus.Infof("failed to close connection: %s", e)
		}
	}()

	// without context timeout
	for i := 0; i < 5; i++ {
		go func() {

		}()
	}

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
	logrus.Errorf("protos error: %+v", err)
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
	go server.StartRPCServer()
	time.Sleep(1 * time.Second)
}

// The TearDownSuite method will be run after all tests have been run.
func (s *ClientTestSuite) TearDownSuite() {
}

func TestClientRequests(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
