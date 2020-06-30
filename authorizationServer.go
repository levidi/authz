package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	opa "github.com/levidi/authz/opa"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
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
			Code: int32(rpc.PERMISSION_DENIED),
		},
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Unauthorized,
				},
				Body: message,
			},
		},
	}
}

func (a *authorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	log.Println(">>> Authorization called check()")

	headers := req.Attributes.Request.Http.Headers
	log.Println(reflect.TypeOf(req.Attributes.Request.Http.Headers))
	log.Println(headers["authorization"])

	body := req.Attributes.Request.Http.Body
	log.Println(reflect.TypeOf(body))
	log.Println(body)

	payload := make(map[string]interface{})
	errUn := json.Unmarshal([]byte(body), &payload)
	if errUn != nil {
		fmt.Println(errUn)
	}
	log.Println("PAYLOAD")
	log.Println(payload)

	b, err := json.Marshal(req.Attributes.Request.Http.Headers)
	if err == nil {
		log.Println("Inbound Headers: ")
		log.Println((string(b)))
	}

	if opa.Eval(payload) {
		log.Println("allow")
		return &auth.CheckResponse{
			Status: &rpcstatus.Status{
				Code: int32(rpc.OK),
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
	log.Println("NOT allow")
	return returnPermissionDenied("Permission Denied"), nil
}
