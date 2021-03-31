# Tipon - Tips gRPC Server
The tips service gRPC server used for internal communication.

## Requirements
You must have the following in order to compile the protocol buffers:
- protoc (https://grpc.io/docs/protoc-installation)
- protoc-gen-go (https://google.golang.org/protobuf/cmd/protoc-gen-go)
- protoc-gen-go-grpc (https://google.golang.org/grpc/cmd/protoc-gen-go-grpc)

## Protocol Buffers
Generate the server/client Go code using the protocol buffer compiler script.
```
/bin/bash protoc
```

## Build & Serve
Build the binary
```
go build
```

Serve the tips gRPC server
```
./tips
```
