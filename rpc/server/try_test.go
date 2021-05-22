package server

import (
	"context"
	"fmt"
	"github.com/cao7113/hellogolang/proto/gosdk/proto/try/v1"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func (s *TryTestSuite) TestTry() {
	asyncStartRPC()

	ct := context.Background()
	conn, err := grpc.DialContext(ct, *ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer conn.Close()

	c := tryv1.NewTryServiceClient(conn)
	req := &tryv1.TryRequest{
		Name:  "Geek",
		Score: uint32(99),
		Gender: &tryv1.TryRequest_Male{
			Male: true,
		},
	}
	logrus.Infof("[client] requesting with %+v", req)
	r, err := c.Try(ct, req)
	if err != nil {
		st := status.Convert(err)
		// codes.Unavailable if rpc shutdown or network conn broken
		logrus.Errorf("error: code: %v message: %s", st.Code(), st.Message())
	}
	if r != nil {
		logrus.Infof("client got msg: %s", r.Message)
	}
}

func (s *TryTestSuite) TestFatal() {
	ct := context.Background()
	conn, err := grpc.DialContext(ct, *ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer conn.Close()
	c := tryv1.NewTryServiceClient(conn)

	req := &tryv1.FatalRequest{}
	logrus.Infof("[client] requesting with %+v", req)
	r, err := c.Fatal(ct, req)
	if err != nil {
		st := status.Convert(err)
		// codes.Unavailable if rpc shutdown or network conn broken
		logrus.Errorf("error: code: %v message: %s", st.Code(), st.Message())
	}
	if r != nil {
		logrus.Infof("client got response: %+v", r)
	}
}

func (s *TryTestSuite) TestSlow() {
	ct := context.Background()
	conn, err := grpc.DialContext(ct, *ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer conn.Close()

	c := tryv1.NewTryServiceClient(conn)
	req := &tryv1.SlowRequest{
		Seconds: 10,
	}
	logrus.Infof("[client] requesting with %+v", req)
	r, err := c.Slow(ct, req)
	if err != nil {
		st := status.Convert(err)
		// codes.Unavailable if rpc shutdown or network conn broken
		logrus.Errorf("error: code: %d message: %s", st.Code(), st.Message())
	}
	if r != nil {
		logrus.Infof("client got msg: %s", r.Msg)
	}
}

func (s *TryTestSuite) TestTryTimeout() {
	ct := context.Background()
	conn, err := grpc.DialContext(ct, *ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer func() {
		if e := conn.Close(); e != nil {
			logrus.Infof("failed to close connection: %s", e)
		}
	}()

	c := tryv1.NewTryServiceClient(conn)
	req := &tryv1.TimeoutRequest{
		TimeoutInMs: 210,
	}
	logrus.Infof("[client] requesting with %+v", req)

	reqCt, cancel := context.WithTimeout(ct, time.Duration(req.TimeoutInMs)*time.Millisecond)
	go func() {
		time.Sleep(350 * time.Millisecond)
		logrus.Infof("client cancel request")
		cancel()
	}()
	r, err := c.Timeout(reqCt, req)
	if err != nil {
		st := status.Convert(err)
		// codes.Unavailable if rpc shutdown or network conn broken
		logrus.Errorf("error: code: %d message: %s", st.Code(), st.Message())
	}
	if r != nil {
		logrus.Infof("client got msg: %s", r.Msg)
	}
}

// call with timeout context

func (s *TryTestSuite) TestWithCallTimeout() {
	ct := context.Background()
	bct, cancel := context.WithTimeout(ct, 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(bct, *ConnAddress, grpc.WithInsecure())
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
				c := tryv1.NewTryServiceClient(conn)
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &tryv1.ContextRequest{
					From: from,
				}
				logrus.Infof("client request with from %s", from)
				r, err := c.Context(ctx, req)
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

func (s *TryTestSuite) TestWithTimeout() {
	ct := context.Background()
	bct, cancel := context.WithTimeout(ct, 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(bct, *ConnAddress, grpc.WithInsecure())
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
				c := tryv1.NewTryServiceClient(conn)
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &tryv1.ContextRequest{
					From: from,
				}
				logrus.Infof("client request with from %s", from)
				r, err := c.Context(ctx, req)
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

func (s *TryTestSuite) TestWithTimeout1() {
	ct := context.Background()
	bct, cancel := context.WithTimeout(ct, 2*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(bct, *ConnAddress, grpc.WithInsecure())
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
	c := tryv1.NewTryServiceClient(conn)
	for i := 0; i < 1; i++ {
		for j := 0; j < 1; j++ {
			//time.Sleep(1 * time.Millisecond)
			go func(i, j int) {
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &tryv1.ContextRequest{
					From: from,
				}
				logrus.Infof("client request with from %s", from)
				r, err := c.Context(ctx, req)
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

func (s *TryTestSuite) TestAllWithoutTimeout() {
	conn, err := grpc.Dial(*ConnAddress, grpc.WithInsecure())
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
				c := tryv1.NewTryServiceClient(conn)
				from := fmt.Sprintf("sub-%d-%d", i, j)
				req := &tryv1.ContextRequest{
					From: from,
				}
				logrus.Infof("client request with from %s", from)
				r, err := c.Context(ctx, req)
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

//func (s *TryTestSuite) TestDetailError() {
//	conn, err := grpc.Dial(*ConnAddress, grpc.WithInsecure())
//	if err != nil {
//		logrus.Fatalf("connect server %v", err)
//	}
//	defer func() {
//		if e := conn.Close(); e != nil {
//			logrus.Infof("failed to close connection: %s", e)
//		}
//	}()
//
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//
//	c := tryv1.NewTryServiceClient(conn)
//	req := &tryv1.HelloRequest{
//		Name:  "try",
//		Error: "error",
//	}
//	r, err := c.Hello(ctx, req)
//	s.Nil(r)
//	s.NotNil(err)
//	se := status.Convert(err)
//	for _, d := range se.Details() {
//		switch info := d.(type) {
//		case *tryv1.Error:
//			logrus.Infof("hit mock error: %+v", info)
//		default:
//			logrus.Fatalf("Unexpected type: %s", info)
//		}
//	}
//	logrus.Errorf("proto error: %+v", err)
//}

func TestTryTestSuite(t *testing.T) {
	suite.Run(t, new(TryTestSuite))
}

type TryTestSuite struct {
	suite.Suite
}

// The SetupSuite method will be run before any tests are run.
func (s *TryTestSuite) SetupSuite() {
}

// The TearDownSuite method will be run after all tests have been run.
func (s *TryTestSuite) TearDownSuite() {
}

func asyncStartRPC() {
	go StartRPCServer()
	time.Sleep(20 * time.Millisecond)
}
