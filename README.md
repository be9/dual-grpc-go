# gRPC-Go dual client PoC

## Intro

This repo contains a primitive gRPC server with 2 methods and a dual client PoC.

The dual client dispatches connections to either a grpc-go 1.53.0 client, or to
[this client](https://github.com/be9/grpc-go/tree/v1.54.0-be9), forked from grpc-go 1.54.0

## Run

1. Run `make server`.
2. Run `make client` in a different console.
3. Check the server output:
   ```
   2023/04/25 15:36:22 gRPC server listening at [::]:50051
   2023/04/25 15:36:22 Starting HTTP at :50052
    
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.53.0]
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.54.0]
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.53.0]
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.54.0]
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.53.0]
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.53.0]
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.54.0]
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.54.0]
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.53.0]
   2023/04/25 15:36:28 GetFoos [agent grpc-go/1.53.0]
   ```
   
## Nuances

1. The fork needs to be adapted
   * Create a branch based on the upstream `go-grpc` release.
   * Rename the module in `go.mod`.
   * `go install github.com/sirkon/go-imports-rename@latest`
   * `go-imports-rename --save 'google.golang.org/grpc => github.com/be9/grpc-go'`
   * Disable duplicate proto registration (`binarylog/grpc_binarylog_v1/binarylog.pb.go`).
2. The dual client cannot pass `opts` to the second client because the type is different.