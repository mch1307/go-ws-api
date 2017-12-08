package main

import (
	"context"
	"fmt"

	"github.com/mch1307/go-ws-api/pb"
	"google.golang.org/grpc"
)

var empty pb.Empty

func main() {
	grpcServer := "localhost:8082"
	conn, err := grpc.Dial(grpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting grpc: ", err)
	}
	defer conn.Close()
	client := pb.NewDeviceServiceClient(conn)
	devices, err := client.GetAllDevices(context.Background(), &empty)
	if err != nil {
		fmt.Println("error creating grpc cli:", err)
	}
	for _, dev := range devices.Device {
		fmt.Println(dev)
	}
}
