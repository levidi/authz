package main

import (
	"log"
	"net"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("listening on %s", listener.Addr())

	grpcServer := grpc.NewServer()
	authServer := &authorizationServer{}
	auth.RegisterAuthorizationServer(grpcServer, authServer)
	healthpb.RegisterHealthServer(grpcServer, &healthServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
