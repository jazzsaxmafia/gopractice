package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "productinfo/proto/productinfopb"
	serverutils "productinfo/server/serverutils"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	log.Print("server listening at port ", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &serverutils.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
