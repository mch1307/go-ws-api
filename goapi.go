package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/mch1307/go-ws-api/db"
	"github.com/mch1307/go-ws-api/pb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

const grpcPort = ":8082"
const httpPort = ":8081"

// runHTTP runs the HTTP server for Rest API proxy
func runHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterDeviceServiceHandlerFromEndpoint(ctx, mux, "localhost"+grpcPort, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpPort, mux)
}

func main() {
	log.Println("Starting application")
	db.InitDB()
	fmt.Println(db.GetAllDevices())
	// start listening for grpc
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}
	// Instanciate new gRPC server
	server := grpc.NewServer()
	//register service
	pb.RegisterDeviceServiceServer(server, new(DeviceService))
	//pb.RegisterNHCServer(server, new(NHCService))

	log.Println("Starting grpc server on port " + grpcPort)

	// Start the gRPC server in goroutine
	go server.Serve(listen)

	// Start the HTTP server for Rest
	log.Println("Starting HTTP server on port " + httpPort)
	runHTTP()
}

type DeviceService struct{}

func (s *DeviceService) GetAllDevices(ctx context.Context, req *pb.Empty) (*pb.Devices, error) {
	devices := db.GetAllDevices()
	return &devices, nil
}

func (s *DeviceService) GetDeviceByID(ctx context.Context, id *pb.ID) (*pb.Device, error) {
	device := db.GetDeviceByID(id.Id)
	return device, nil
}

func (s *DeviceService) SwitchDevice(ctx context.Context, device *pb.UpdateDevice) (*pb.Device, error) {
	updatedDevice, err := db.SwitchDevice(device.Id, device.Value)
	if err != nil {
		log.Println("error updating device ", err)
	}
	return updatedDevice, err
}

func (s *DeviceService) RegisterDevice(ctx context.Context, device *pb.Device) (*pb.Device, error) {
	return nil, nil
}
