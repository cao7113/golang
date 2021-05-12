package server

import (
	"context"
	"fmt"
	pingv1 "github.com/cao7113/hellogolang/proto/gosdk/proto/ping/v1"
	"time"
)

type PingServer struct {
	pingv1.UnimplementedPingServiceServer
}

func (p PingServer) Ping(ctx context.Context, req *pingv1.PingRequest) (*pingv1.PingResponse, error) {
	msg := fmt.Sprintf("pong from %s", req.From)
	resp := &pingv1.PingResponse{
		Message:   msg,
		Timestamp: time.Now().Unix(),
	}
	return resp, nil
}
