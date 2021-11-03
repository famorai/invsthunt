package main

import (
	"log"
	"net"

	"github.com/famorai/training/pb"
	"github.com/famorai/training/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Coud not to be connect:%v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTokenServiceServer(grpcServer, services.NewTokenService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Coud not be server:%v", err)
	}

}
