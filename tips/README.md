# Tipon - Tips gRPC Server
The tips service gRPC server used for internal communication.

## Protocol Buffers
Generate the server/client Go code using the protocol buffer compiler script (must have `protoc` installed).
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
