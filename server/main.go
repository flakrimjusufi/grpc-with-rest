package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	db "server/main.go/database"
	userpb "server/main.go/proto"
)

var database = db.Connect()

type User struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
}

type ListAllUsers struct {
	User []*User
}

type userServer struct {
	userpb.UnimplementedUserServiceServer
}

func (as *userServer) CreateUser(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	user := User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}

	database.NewRecord(user)
	database.Debug().Create(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) UpdateUser(ctx context.Context, in *userpb.User) (*userpb.User, error) {

	name := in.GetName()
	email := in.GetEmail()
	phoneNumber := in.GetPhoneNumber()

	var user User
	database.Where("name =?", name).Find(&user)

	user.Email = email
	user.PhoneNumber = phoneNumber

	database.Debug().Save(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) DeleteUser(ctx context.Context, in *userpb.User) (*userpb.Message, error) {
	name := in.GetName()
	var user User
	database.Where("name =?", name).Find(&user)
	database.Debug().Delete(&user)

	return &userpb.Message{Message: user.Name + " Deleted successfully!"}, nil
}

func (as *userServer) ListUsers(ctx context.Context, in *userpb.User) (*userpb.ListUser, error) {

	list := make([]*userpb.User, 0)
	database.Debug().Table("users").Find(&list)
	fmt.Println("{}", list)

	return &userpb.ListUser{
		Users: list,
	}, nil
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

	//Auto-Migration of User Model
	database.AutoMigrate(&User{})
	defer database.Close()

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
