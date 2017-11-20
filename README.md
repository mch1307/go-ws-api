# go-ws-api

Sample gRPC / Rest API project in Go

Uses [protoc](http://github.com/google/protobuf), [gRPC](http://grpc.io) and [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) for Rest and OpenAPI/Swagger

## Context

In this example we have a small "home control" system that simply turn on/off some lights.

As our system grows in popularity, we want to open it to the world, and to achieve this, we will develop an API. For the first mileston, we want to expose the following endpoints:

* list all registered devices
* find a device by it's id
* switch a device on, off or dim
* create / register a new device

The following is a device's JSON representation we use in this project

```json
{
    "id": 1,
    "hardware": "philips",
    "name": "light",
    "location" : "kitchen",
    "type": "onOff",
    "state": 100
}
```

## How it was built

* write the [protoc spec](./pb/device.proto)
* generate [pb services and message types file](./pb/device.pb.go)

    ```cmd
    protoc -I. -I%GOPATH%\src -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --go_out=plugins=grpc:. device.proto
    ```

* implement the code
* generate the [grpc-gateway file](./pb/device.pb.gw.go)

    ```cmd
    protoc -I. -I%GOPATH%\src -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --grpc-gateway_out=logtostderr=true:. device.proto
    ```

* generate the [OpenAPI spec/doc](./device.swagger.json)

    ```cmd
    protoc -I. -I%GOPATH%\src -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis --swagger_out=logtostderr=true:. device.proto
    ```
