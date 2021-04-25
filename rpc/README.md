# gRPC

## Quick start
https://grpc.io/docs/languages/go/quickstart/

## Examples
https://github.com/grpc/grpc-go/tree/master/examples

## buf CLI
https://docs.buf.build/tour-1

### Googleapis examples

https://github.com/googleapis/googleapis

```
GOOGLEAPIS_COMMIT="d4aa417ed2bba89c2d216900282bddfdafef6128"
git clone https://github.com/googleapis/googleapis
cd googleapis
git reset --hard "${GOOGLEAPIS_COMMIT}"
```

## Todo

rpc 版本问题解决方案！！！

//go:generate protoc --go_out=plugins=grpc:. protos/*.proto

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
protos/hello.proto

export GO111MODULE=on  # Enable module mode
go get google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc