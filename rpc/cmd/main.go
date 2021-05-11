package main

import (
	"github.com/cao7113/hellogolang/rpc/server"
)

func main() {
	//tracer.Start(
	//	tracer.WithEnv("testing"),
	//	tracer.WithService(config.Config.ServiceName),
	//	tracer.WithServiceVersion("v0.0.1"),
	//)
	//defer tracer.Stop()

	server.StartRPCServer()
}
