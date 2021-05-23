package main

import (
	"flag"
	"github.com/cao7113/hellogolang/rpc/server"
)

func main() {
	//tracer.Start(
	//	tracer.WithEnv("testing"),
	//	tracer.WithService(config.Config.ServiceName),
	//	tracer.WithServiceVersion("v0.0.1"),
	//)
	//defer tracer.Stop()

	pPort := flag.Int("port", 50051, "-port=50051")
	flag.Parse()
	port := *pPort
	server.StartRPCServer(port, "")
}
