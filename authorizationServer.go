package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	opa "github.com/levidi/authz/opa"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

type authorizationServer struct{}

func returnUnAuthenticated(message string) *auth.CheckResponse {
	return &auth.CheckResponse{
		Status: &rpcstatus.Status{
			Code: int32(rpc.UNAUTHENTICATED),
		},
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Forbidden,
				},
				Body: message,
			},
		},
	}
}

func returnPermissionDenied(message string) *auth.CheckResponse {
	return &auth.CheckResponse{
		Status: &rpcstatus.Status{
			Code:    int32(rpc.PERMISSION_DENIED),
			Details: []*anypb.Any{{Value: []byte(`"allow": "false"`)}},
		},
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Forbidden,
				},
				Body: message,
			},
		},
	}
}

func (a *authorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {

	log.Println("START")
	log.Println(req)
	log.Println("END")

	body := req.Attributes.Request.Http.Body

	payload := make(map[string]interface{})
	errUn := json.Unmarshal([]byte(body), &payload)
	if errUn != nil {
		fmt.Println(errUn)
	}

	if opa.Eval(payload) {

		return &auth.CheckResponse{
			Status: &rpcstatus.Status{
				Code:    int32(rpc.OK),
				Details: []*anypb.Any{{Value: []byte(`"allow": "true"`)}},
			},
			HttpResponse: &auth.CheckResponse_OkResponse{

				OkResponse: &auth.OkHttpResponse{

					Headers: []*core.HeaderValueOption{
						{
							Header: &core.HeaderValue{
								Key:   "x-custom-header-from-authz",
								Value: "some value",
							},
						},
					},
				},
			},
		}, nil
	}
	return returnPermissionDenied(`"reason": "Permission Denied"`), nil
}
