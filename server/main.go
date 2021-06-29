package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	userpb "server/main.go/proto"
)

type userServer struct {
	userpb.UnimplementedUserServiceServer
}

func (as *userServer) CreateUser(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	return &userpb.User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}, nil
}

func (as *userServer) UpdateUser(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	return &userpb.User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}, nil
}

func (as *userServer) DeleteUser(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	return &userpb.User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}, nil
}

func (as *userServer) ListUsers(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	return &userpb.User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the User service to the server
	userpb.RegisterUserServiceServer(s, &userServer{})

	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = userpb.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}