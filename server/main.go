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
	"os"
	db "server/main.go/database"
	userpb "server/main.go/proto"
)

var database = db.Connect().Debug()

type User struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
}

type userServer struct {
	userpb.UnimplementedUserServiceServer
}

func (as *userServer) SayHello(ctx context.Context, in *userpb.User) (*userpb.Message, error) {
	return &userpb.Message{Message: "Hello " + in.Name}, nil
}

func (as *userServer) CreateUser(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	user := User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}

	database.NewRecord(user)
	database.Create(&user)

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

	database.Save(&user)

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
	database.Where("deleted_at is null").Order("created_at desc").Find(&list)
	return &userpb.ListUser{
		Users: list,
	}, nil
}

func (as *userServer) GetUserByName(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	name := in.GetName()
	var user User
	database.Where(&User{Name: name}).Find(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) GetUserById(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	id := in.GetId()
	var user User
	database.Where("id = ?", id).Find(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
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
		Addr:    fmt.Sprintf(":%s", os.Getenv("port")),
		Handler: gwmux,
	}

	//Auto-Migration of User Model
	database.AutoMigrate(&User{})
	defer database.Close()

	log.Println(fmt.Sprintf("Serving gRPC-Gateway on %s:%s", os.Getenv("host"), os.Getenv("port")))
	log.Fatalln(gwServer.ListenAndServe())
}
