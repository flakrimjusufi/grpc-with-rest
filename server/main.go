package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	db "server/main.go/database"
	models "server/main.go/models"
	userpb "server/main.go/proto"
)

const (
	address     = "localhost:8080"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorWhite  = "\033[37m"
	colorPurple = "\033[35m"
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
)

var database = db.Connect().Debug()

type userServer struct {
	userpb.UnimplementedUserServiceServer
}

type creditCardServer struct {
	userpb.UnimplementedCreditCardServiceServer
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (as *userServer) SayHello(ctx context.Context, in *userpb.User) (*userpb.Message, error) {
	return &userpb.Message{Message: "Hello " + in.Name}, nil
}

func (as *userServer) CreateUser(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	user := models.User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}

	database.NewRecord(user)
	database.Create(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) UpdateUserByName(ctx context.Context, in *userpb.User) (*userpb.User, error) {

	name := in.GetName()
	email := in.GetEmail()
	phoneNumber := in.GetPhoneNumber()

	var user models.User
	database.Where("name =?", name).Find(&user)

	user.Email = email
	user.PhoneNumber = phoneNumber

	database.Save(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) UpdateUserById(ctx context.Context, in *userpb.User) (*userpb.User, error) {

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

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) DeleteUser(ctx context.Context, in *userpb.User) (*userpb.Message, error) {
	name := in.GetName()
	var user models.User
	database.Where("name =?", name).Find(&user)
	database.Delete(&user)

	return &userpb.Message{Message: user.Name + " Deleted successfully!"}, nil
}

func (as *userServer) ListUsers(ctx context.Context, in *userpb.User) (*userpb.ListUser, error) {

	list := make([]*userpb.User, 0)
	database.Where("deleted_at is null").Order("created_at desc").Limit(100).Find(&list)
	return &userpb.ListUser{
		Users: list,
	}, nil
}

func (as *userServer) GetUserByName(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	log.Println(colorPurple, "[userService] - [rpc GetUserByName] -> ", colorBlue, "Received person's name: ", in.GetName())
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	name := in.GetName()
	var user models.User
	database.Where(&models.User{Name: name}).Find(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) GetUserById(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	id := in.GetId()
	var user models.User
	database.Where("id = ?", id).Find(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) FindUserFromGetUserByIdRPC(ctx context.Context, in *userpb.User) (*userpb.User, error) {

	user, _ := as.GetUserById(ctx, in)

	return &userpb.User{Id: user.GetId(), Name: user.GetName(), Email: user.GetEmail(), PhoneNumber: user.GetPhoneNumber()}, nil
}

func (as *userServer) FindUserFromGetUserByNameRPC(ctx context.Context, in *userpb.User) (*userpb.User, error) {

	user, _ := as.GetUserByName(ctx, in)

	return &userpb.User{Id: user.GetId(), Name: user.GetName(), Email: user.GetEmail(), PhoneNumber: user.GetPhoneNumber()}, nil
}

func (fu *creditCardServer) CreditCards(ctx context.Context, in *userpb.CreditCard) (*userpb.ListCreditCards, error) {

	var list []*userpb.CreditCard
	var creditCards []*models.CreditCards
	database.Order("created_at desc").Find(&creditCards)

	for _, card := range creditCards {
		list = append(list, &userpb.CreditCard{
			Id:          uint32(card.ID),
			Name:        card.Name,
			Email:       card.Email,
			PhoneNumber: card.PhoneNumber,
			Address:     card.Address,
			Country:     card.Country,
			City:        card.City,
			Zip:         card.Zip,
			Cvv:         card.CVV,
			CreatedAt:   timestamppb.New(card.CreatedAt),
			UpdatedAt:   timestamppb.New(card.UpdatedAt),
			DeletedAt:   timestamppb.New(card.DeletedAt),
		})
	}

	return &userpb.ListCreditCards{
		CreditCards: list,
	}, nil
}

func (fu *creditCardServer) GetCreditCardByUserName(ctx context.Context, in *userpb.CreditCard) (*userpb.CreditCard, error) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	name := in.GetName()
	var creditCard models.CreditCards
	database.Where(&models.CreditCards{Name: name}).Find(&creditCard)

	log.Println(colorPurple, "[creditCardService] - [rpc GetCreditCardByUserName] -> ", colorGreen, "Now sending the response (credit card of selected user) to client side...")

	return &userpb.CreditCard{Id: uint32(creditCard.ID), Name: creditCard.Name, Email: creditCard.Email, PhoneNumber: creditCard.PhoneNumber,
		Address: creditCard.Address, Country: creditCard.Country, City: creditCard.City, Zip: creditCard.Zip, Cvv: creditCard.CVV,
		CreatedAt: timestamppb.New(creditCard.CreatedAt)}, nil
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
	userpb.RegisterCreditCardServiceServer(s, &creditCardServer{})

	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	maxMsgSize := 1024 * 1024 * 20
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register User Service
	err = userpb.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	newServer := userpb.RegisterCreditCardServiceHandler(context.Background(), gwmux, conn)
	if newServer != nil {
		log.Fatalln("Failed to register gateway:", newServer)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("server_port")),
		Handler: cors(gwmux),
	}

	log.Println(fmt.Sprintf("Serving gRPC-Gateway on %s:%s", os.Getenv("server_host"), os.Getenv("server_port")))
	log.Fatalln(gwServer.ListenAndServe())
}
