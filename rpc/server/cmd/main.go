package main

import (
	"github.com/cao7113/hellogolang/config"
	"github.com/cao7113/hellogolang/rpc/server"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithEnv("testing"),
		tracer.WithService(config.Config.ServiceName),
		tracer.WithServiceVersion("v0.0.1"),
	)
	defer tracer.Stop()

	server.StartRPCServer()
}
