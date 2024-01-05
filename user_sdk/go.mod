module github.com/flutterninja9/shoppie/user_sdk

go 1.21.0

require (
	github.com/flutterninja9/shoppie/user_service/protos v0.0.1
	github.com/google/uuid v1.5.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/grpc v1.60.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/flutterninja9/shoppie/user_service/protos => ../user_service/protos
