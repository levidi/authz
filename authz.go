package main

import (
	"log"
	"net"

	envoy_service_auth_v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"github.com/levidi/authz/policypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("listening on %s", listener.Addr())

	grpcServer := grpc.NewServer()

	envoy_service_auth_v2.RegisterAuthorizationServer(grpcServer, &authorizationServer{})
	grpc_health_v1.RegisterHealthServer(grpcServer, &healthServer{})
	policypb.RegisterPolicyServiceServer(grpcServer, &policyServer{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
