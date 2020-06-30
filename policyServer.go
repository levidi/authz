package main

import (
	"context"
	"fmt"

	"github.com/levidi/authz/policypb"
)

type policyServer struct{}

func (s *policyServer) SetPolicy(ctx context.Context, req *policypb.PolicyRequest) (res *policypb.PolicyResponse, err error) {
	fmt.Printf("Received SetPolicy RPC: %v", req)
	return &policypb.PolicyResponse{Message: "ok", Status: 200}, nil
}
