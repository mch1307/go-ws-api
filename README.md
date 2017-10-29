# go-ws-api

Sample gRPC / Rest API project in Go

Uses [protoc](http://github.com/google/protobuf), [gRPC](grpc.io) and [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) for Rest and OpenAPI/Swagger

## How it was built

* write the [protoc spec](./pb/device.proto)
* generate [pb services and message types file](./pb/device.pb.go)

    `protoc -I. -I%GOPATH%\src -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go_out=plugins=grpc:. device.proto`

* implement the code
* generate the [grpc-gateway file](./pb/device.pb.gw.go)

    `protoc -I. -I%GOPATH%\src -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --grpc-gateway_out=logtostderr=true:. device.proto`

* generate the [OpenAPI spec/doc](./device.swagger.json)

    `protoc -I. -I%GOPATH%\src -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --swagger_out=logtostderr=true:. device.proto`
