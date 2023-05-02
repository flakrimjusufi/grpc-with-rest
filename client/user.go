package client

import (
	"context"
	"fmt"
	"log"

	"github.com/flakrimjusufi/grpc-with-rest/models"
	"github.com/flakrimjusufi/grpc-with-rest/proto"
	userpb "github.com/flakrimjusufi/grpc-with-rest/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

func (as *UserServer) SayHello(ctx context.Context, in *proto.User) (*proto.Message, error) {
	return &proto.Message{Message: "Hello " + in.Name}, nil
}

func (as *UserServer) CreateUser(ctx context.Context, in *proto.User) (*proto.User, error) {
	user := models.User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}

	database.NewRecord(user)
	database.Create(&user)

	return &proto.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *UserServer) UpdateUserByName(ctx context.Context, in *proto.User) (*proto.User, error) {

	name := in.GetName()
	if name == "" {
		return &proto.User{}, status.Error(codes.InvalidArgument, "User's name not specified")
	}
	email := in.GetEmail()
	phoneNumber := in.GetPhoneNumber()

	var user models.User
	database.Where("name =?", name).Find(&user)

	user.Email = email
	user.PhoneNumber = phoneNumber

	database.Save(&user)

	return &proto.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *UserServer) UpdateUserByID(ctx context.Context, in *proto.User) (*proto.User, error) {

	id := in.GetId()
	name := in.GetName()
	email := in.GetEmail()
	phoneNumber := in.GetPhoneNumber()

	var user models.User
	database.Where("id =?", id).Find(&user)

	user.Name = name
	user.Email = email
	user.PhoneNumber = phoneNumber

	database.Save(&user)

	return &proto.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *UserServer) DeleteUser(ctx context.Context, in *proto.User) (*proto.Message, error) {
	name := in.GetName()
	if name == "" {
		return &proto.Message{}, status.Error(codes.InvalidArgument, "User's name not specified")
	}
	var user models.User
	rowsAffected := database.Where("name =?", name).Find(&user).RowsAffected

	if rowsAffected == 0 {
		return &proto.Message{}, status.Error(codes.NotFound, "Cannot find a User with this name!")
	}
	database.Delete(&user)

	return &proto.Message{Message: user.Name + " Deleted successfully!"}, nil
}

func (as *UserServer) ListUsers(ctx context.Context, in *proto.User) (*proto.ListUser, error) {

	list := make([]*proto.User, 0)
	database.Where("deleted_at is null").Order("created_at desc").Limit(100).Find(&list)
	return &proto.ListUser{
		Users: list,
	}, nil
}

func (as *UserServer) GetUserByName(ctx context.Context, in *proto.User) (*proto.User, error) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	log.Println(colorPurple, "[userService] - [rpc GetUserByName] -> ",
		colorBlue, "Received person's name: ", in.GetName())
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	name := in.GetName()
	if name == "" {
		return &proto.User{}, status.Error(codes.InvalidArgument, "User's name not specified")
	}
	var user models.User
	database.Where(&models.User{Name: name}).Find(&user)

	return &proto.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *UserServer) GetUserByID(ctx context.Context, in *proto.User) (*proto.User, error) {
	id := in.GetId()
	var user models.User
	rowsAffected := database.Where("id = ?", id).Find(&user).RowsAffected

	if rowsAffected == 0 {
		return &proto.User{}, status.Error(codes.NotFound, "Cannot find a User with this id!")
	}

	return &proto.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}
