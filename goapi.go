package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"sgithub.sgbt.lu/champam1/go-api/db"
	"sgithub.sgbt.lu/champam1/go-api/pb"
)

const grpcPort = ":8082"
const httpPort = ":8081"

func runHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterNHCHandlerFromEndpoint(ctx, mux, "localhost"+grpcPort, opts)
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

	//opts := []grpc.DialOption{grpc.WithInsecure()}
	server := grpc.NewServer()
	pb.RegisterNHCServer(server, new(NHCService))
	log.Println("Starting grpc server on port " + grpcPort)
	go server.Serve(listen)
	log.Println("Starting HTTP server on port " + httpPort)
	runHTTP()
}

type NHCService struct{}

func (s *NHCService) GetAllDevices(ctx context.Context, req *pb.Empty) (*pb.Devices, error) {
	devices := db.GetAllDevices()
	return &devices, nil
}

func (s *NHCService) GetDeviceByID(ctx context.Context, id *pb.ID) (*pb.Device, error) {
	device := db.GetDeviceByID(id.Id)
	return device, nil
}
