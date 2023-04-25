module github.com/be9/dual-grpc-go

go 1.18

require (
	github.com/be9/grpc-go v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.8.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.53.0

replace github.com/be9/grpc-go => github.com/be9/grpc-go v0.0.0-20230425082119-2a933b55cbca
