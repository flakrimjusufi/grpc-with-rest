package client

import (
	"context"
	"fmt"
	"log"

	"github.com/flakrimjusufi/grpc-with-rest/models"
	userpb "github.com/flakrimjusufi/grpc-with-rest/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
)

// UserServer - the grpc server of users
type UserServer struct {
	userpb.UnimplementedUserServiceServer
}

// SayHello - the service that prints a given name in the output
func (us *UserServer) SayHello(ctx context.Context, in *userpb.User) (*userpb.Message, error) {
	return &userpb.Message{Message: "Hello " + in.Name}, nil
}

// CreateUser - the service that creates a user in the Users table and returns userpb.User
func (us *UserServer) CreateUser(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	user := models.User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}

	database.NewRecord(user)
	result := database.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "UserServer - CreateUser: %v", result.Error)
	}

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

// UpdateUserByName - the service that updates a user by its name in the Users table and returns userpb.User
func (us *UserServer) UpdateUserByName(ctx context.Context, in *userpb.User) (*userpb.User, error) {

	name := in.GetName()
	if name == "" {
		return &userpb.User{}, status.Error(codes.InvalidArgument, "User's name not specified")
	}
	email := in.GetEmail()
	phoneNumber := in.GetPhoneNumber()

	var user models.User
	result := database.Where("name =?", name).Find(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "UserServer - UpdateUserByName: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "UserServer - UpdateUserByName - No user was found with the given name")
	}

	user.Email = email
	user.PhoneNumber = phoneNumber

	database.Save(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

// UpdateUserByID - the service that updates a user by its ID in the Users table and returns userpb.User
func (us *UserServer) UpdateUserByID(ctx context.Context, in *userpb.User) (*userpb.User, error) {

	id := in.GetId()
	name := in.GetName()
	email := in.GetEmail()
	phoneNumber := in.GetPhoneNumber()

	var user models.User
	result := database.Where("id =?", id).Find(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "UserServer - UpdateUserByID: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "UserServer - UpdateUserByID - No user was found with the id")
	}

	user.Name = name
	user.Email = email
	user.PhoneNumber = phoneNumber

	database.Save(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

// DeleteUser - the service that deletes a user by its name in the Users table and returns userpb.User
func (us *UserServer) DeleteUser(ctx context.Context, in *userpb.User) (*userpb.Message, error) {
	name := in.GetName()
	if name == "" {
		return &userpb.Message{}, status.Error(codes.InvalidArgument, "User's name not specified")
	}
	var user models.User
	rowsAffected := database.Where("name =?", name).Find(&user).RowsAffected

	if rowsAffected == 0 {
		return &userpb.Message{}, status.Error(codes.NotFound, "No user was found with the given name")
	}
	database.Delete(&user)

	return &userpb.Message{Message: user.Name + " Deleted successfully!"}, nil
}

// ListUsers - the service that lists the users in the Users table and returns userpb.User
func (us *UserServer) ListUsers(ctx context.Context, in *userpb.User) (*userpb.ListUser, error) {

	list := make([]*userpb.User, 0)
	result := database.Where("deleted_at is null").Order("created_at desc").Limit(100).Find(&list)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "UserServer - ListUsers: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "UserServer - ListUsers - No users yet.")
	}
	return &userpb.ListUser{
		Users: list,
	}, nil
}

// GetUserByName - the service that gets the user by name in the Users table and returns userpb.User
func (us *UserServer) GetUserByName(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	log.Println(colorPurple, "[userService] - [rpc GetUserByName] -> ",
		colorBlue, "Received person's name: ", in.GetName())
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	name := in.GetName()
	if name == "" {
		return &userpb.User{}, status.Error(codes.InvalidArgument, "User's name not specified")
	}
	var user models.User
	result := database.Where(&models.User{Name: name}).Find(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "UserServer - GetUserByName: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "UserServer - GetUserByName - No user was found with the given name.")
	}

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

// GetUserByID - the service that gets the user by ID in the Users table and returns userpb.User
func (us *UserServer) GetUserByID(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	id := in.GetId()
	var user models.User
	result := database.Where("id = ?", id).Find(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "UserServer - GetUserByID: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return &userpb.User{}, status.Error(codes.NotFound, "Cannot find a User with this id!")
	}

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}
