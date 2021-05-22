package server

import (
	"context"
	"errors"
	streamv1 "github.com/cao7113/hellogolang/proto/gosdk/proto/stream/v1"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"testing"
	"time"
)

func (s *StreamTestSuite) TestHiStream() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, *ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer conn.Close()

	c := streamv1.NewStreamServiceClient(conn)
	req := &streamv1.HiRequest{
		From:     "testing",
		MsgCount: 10,
	}
	logrus.Infof("[client] requesting with %+v", req)

	rs, err := c.Hi(ctx, req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			logrus.Fatalf("error: code: %d message: %s", st.Code(), st.Message())
		} else {
			logrus.Fatalf("Hi error: %v", err)
		}
	}

	if rs != nil {
		cnt := 0
		for {
			hr, err := rs.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					logrus.Infof("got %d messages, then server stream end", cnt)
				} else {
					logrus.Errorf("Recv() unexpected error: %v", err)
				}
				break
			}
			if hr != nil {
				logrus.Infof("hi idx=%d response: %+v", hr.Index, hr)
			}
			cnt++
			//time.Sleep(10 * time.Millisecond)
		}
		s.EqualValues(req.MsgCount, cnt)
	}
}

func (s *StreamTestSuite) TestClientStream() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, *ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer conn.Close()

	c := streamv1.NewStreamServiceClient(conn)
	cs, err := c.ClientStream(ctx)
	if err != nil {
		logrus.Fatalf("ClientStream fatal: %s", err.Error())
	}
	cnt := 3
	for i := 0; i < cnt; i++ {
		req := &streamv1.ClientStreamRequest{
			From:  "testing",
			Index: int32(i),
		}
		err = cs.Send(req)
		if err != nil {
			logrus.Fatalf("Send() error: %s", err.Error())
		}
		logrus.Infof("[client] sent msg: %+v", req)
	}
	time.Sleep(2 * time.Second)
}

func (s *StreamTestSuite) TestBiStream() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, *ConnAddress, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("connect server %v", err)
	}
	defer conn.Close()

	c := streamv1.NewStreamServiceClient(conn)
	bs, err := c.BiStream(ctx)
	if err != nil {
		logrus.Fatalf("ClientStream fatal: %s", err.Error())
	}

	// push messages
	go func() {
		cnt := 7
		for i := 0; i < cnt; i++ {
			req := &streamv1.BiStreamRequest{
				From:  "testing",
				Index: int32(i),
			}
			err = bs.Send(req)
			if err != nil {
				logrus.Fatalf("Send() error: %s", err.Error())
			}
			logrus.Infof("[client] sent msg: %+v", req)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// get messages
	go func() {
		for {
			bsResp, err := bs.Recv()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					logrus.Errorf("[client] Recv() error: %s", err.Error())
				}
				break
			}
			logrus.Infof("[client] got response: %+v", bsResp)
		}
	}()

	time.Sleep(3 * time.Second)
	//select {}
}

func TestStreamTestSuite(t *testing.T) {
	suite.Run(t, new(StreamTestSuite))
}

type StreamTestSuite struct {
	suite.Suite
}

// The SetupSuite method will be run before any tests are run.
func (s *StreamTestSuite) SetupSuite() {
	go StartRPCServer()
	//time.Sleep(1 * time.Second)
}

// The TearDownSuite method will be run after all tests have been run.
func (s *StreamTestSuite) TearDownSuite() {
}
