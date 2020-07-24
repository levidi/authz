package main

import (
	"context"
	"log"
	"os"
	"testing"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"google.golang.org/grpc"
)

var clientAuth auth.AuthorizationClient

func TestMain(m *testing.M) {

	opts := grpc.WithInsecure()
	clientConn, err := grpc.Dial("0.0.0.0:50051", opts)
	if err != nil {
		log.Fatalf("could not connecti: %v", err)
	}

	clientAuth = auth.NewAuthorizationClient(clientConn)
	code := m.Run()
	os.Exit(code)
}

func TestCheck(t *testing.T) {
	t.Parallel()
	headers := make(map[string]string)
	headers["authorization"] = "levi"
	body := `{"name": "Levi Di Tomazzo", "age": 34, "profile": "administrator"}`

	log.Println(body)
	checkReq := formatCheckRequest(headers, body)
	res, err := clientAuth.Check(context.Background(), checkReq)
	if err != nil {
		t.Errorf("Error while checking server auth: %v", err)
	}
	t.Logf("Auth server: %v", res)
}

func formatCheckRequest(headers map[string]string, body string) (checkReq *auth.CheckRequest) {
	attContextHTTPReq := &auth.AttributeContext_HttpRequest{Headers: headers, Body: body}
	attContext := &auth.AttributeContext{Request: &auth.AttributeContext_Request{Http: attContextHTTPReq}}
	return &auth.CheckRequest{Attributes: attContext}
}
