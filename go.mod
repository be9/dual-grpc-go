module github.com/be9/dual-grpc-go

go 1.18

require (
	golang.org/x/net v0.8.0
	google.golang.org/grpc v1.53.0
	google.golang.org/grpc/v2 v2.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.53.0

replace google.golang.org/grpc/v2 => google.golang.org/grpc v1.54.0
