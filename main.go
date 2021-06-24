package main

import (
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	gw "grpc-with-rest/proto"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux, err := newGateway(ctx)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	log.Println("grpc-gateway listen on localhost:8080")
	return http.ListenAndServe(":8080", mux)
}

func newGateway(ctx context.Context) (http.Handler, error) {

	opts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	if err := gw.RegisterGreeterHandlerFromEndpoint(ctx, gwmux, ":9192", opts); err != nil {
		return nil, err
	}

	return gwmux, nil
}

func main() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}