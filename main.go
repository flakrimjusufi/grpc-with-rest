package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/flakrimjusufi/grpc-with-rest/client"
	db "github.com/flakrimjusufi/grpc-with-rest/database"
	"github.com/flakrimjusufi/grpc-with-rest/helper"
	userpb "github.com/flakrimjusufi/grpc-with-rest/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var grpcServerPort = fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("GRPC_SERVER_PORT"))
var grpcGatewayPort = fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("GRPC_GATEWAY_SERVER_PORT"))

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

var idleTimeout = 5 * time.Second     // the amount of time we are willing to keep a call idle
var contextTimeout = 10 * time.Second // the amount of time we are willing to wait for a call to be completed

func main() {
	if os.Getenv("GRPC_SERVER_PORT") == "" {
		e := godotenv.Load() //Load .env file for local environment
		if e != nil {
			fmt.Println(e)
		}
	}
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_SERVER_PORT")))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed to connect to the database: %v", err))
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Create service instances with the DB connection
	userServer := &client.UserServer{DB: dbConn.Conn}
	creditCardServer := &client.CreditCardServer{DB: dbConn.Conn}

	// Attach the services to the server
	userpb.RegisterUserServiceServer(s, userServer)
	userpb.RegisterCreditCardServiceServer(s, creditCardServer)

	// Serve gRPC server
	log.Printf("Serving gRPC on %s:%s", os.Getenv("SERVER_HOST"), os.Getenv("GRPC_SERVER_PORT"))
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// All the calls that exceeds the 10 seconds threshold, will be cancelled by the server
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	// max message size the gRPC clients can process
	maxMsgSize := 1024 * 1024 * 20

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		grpcServerPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithReadBufferSize(maxMsgSize),
		grpc.WithWriteBufferSize(maxMsgSize),
		grpc.WithKeepaliveParams(kacp),
		grpc.WithIdleTimeout(idleTimeout),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register User Service
	err = userpb.RegisterUserServiceHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	newServer := userpb.RegisterCreditCardServiceHandler(ctx, gwmux, conn)
	if newServer != nil {
		log.Fatalln("Failed to register gateway:", newServer)
	}

	gwServer := &http.Server{
		Addr:    grpcGatewayPort,
		Handler: helper.Cors(gwmux),
	}

	log.Printf("Serving gRPC-Gateway on %s:%s", os.Getenv("SERVER_HOST"),
		os.Getenv("GRPC_GATEWAY_SERVER_PORT"))
	log.Fatalln(gwServer.ListenAndServe())
}
