package server

import (
	"errors"
	"fmt"
	streamv1 "github.com/cao7113/hellogolang/proto/gosdk/proto/stream/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"sync"
)

type StreamServer struct {
	streamv1.UnimplementedStreamServiceServer
}

func (s StreamServer) Hi(req *streamv1.HiRequest, svr streamv1.StreamService_HiServer) error {
	logrus.Infof("[server] handling request with: %+v", req)
	for i := int32(0); i < req.MsgCount; i++ {
		hr := &streamv1.HiResponse{
			Message: fmt.Sprintf("index=%d response for %s request", i, req.From),
			Index:   i,
		}
		err := svr.Send(hr)
		if err != nil {
			return err
		}
		logrus.Infof("[server] sent %d msg: %+v", i, hr)
	}
	return nil
}

func (s StreamServer) ClientStream(cs streamv1.StreamService_ClientStreamServer) error {
	cnt := 0
	for {
		csr, err := cs.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				logrus.Errorf("Recv() unexecpted error: %s", err.Error())
				return status.Errorf(codes.Internal, "Recv() error: %s", err.Error())
			}
			break
		}
		logrus.Infof("[server] got %d from idx=%d request: %+v", cnt, csr.Index, csr)
		cnt++
	}
	return nil
}

func (s StreamServer) BiStream(svr streamv1.StreamService_BiStreamServer) error {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// client stream, pull messages
	go func() {
		cnt := 0
		for {
			csr, err := svr.Recv()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					logrus.Errorf("[server] Recv() unexecpted error: %s", err.Error())
				}
				break
			}
			logrus.Infof("[server] got %d from idx=%d request: %+v", cnt, csr.Index, csr)
			cnt++
		}
		wg.Done()
	}()

	// server stream, push messages
	go func() {
		// todo 与收到的消息联动起来 更有趣
		for i := int32(0); i < 3; i++ {
			hr := &streamv1.BiStreamResponse{
				Message: fmt.Sprintf("index=%d response message", i),
				Index:   i,
			}
			err := svr.Send(hr)
			if err != nil {
				logrus.Errorf("[server] Send() unexpected error: %s", err.Error())
				break
			}
			logrus.Infof("[server] sent %d msg: %+v", i, hr)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}
