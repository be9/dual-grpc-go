.PHONY: proto
proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/testservice/testservice.proto

.PHONY: server
server:
	@go run server/main.go

.PHONY: client
client:
	@go run client/main.go --host localhost:50051

