package main

import (
	"app/api"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"net/http"
)

type helloWorldServer struct{}

func (h *helloWorldServer) SayHello(ctx context.Context, in *api.HelloIn) (*api.HelloOut, error) {
	return &api.HelloOut{
		Out: fmt.Sprintf("Hello %s", in.In),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterHelloWorldServer(grpcServer, &helloWorldServer{})
	go grpcServer.Serve(listener)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = api.RegisterHelloWorldHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		fmt.Println("Error occured: ", err)
	}

	http.ListenAndServe(":8080", mux)
}
