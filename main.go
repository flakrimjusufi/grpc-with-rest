package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	db "github.com/flakrimjusufi/grpc-with-rest/database"
	"github.com/flakrimjusufi/grpc-with-rest/helper"
	models "github.com/flakrimjusufi/grpc-with-rest/models"
	userpb "github.com/flakrimjusufi/grpc-with-rest/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
)

const (
	colorGreen  = "\033[32m"
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

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if helper.AllowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
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

func (as *userServer) CreatePayload(ctx context.Context, in *userpb.Payload) (*userpb.Payload, error) {

	list := make([]interface{}, 0)

	for _, attributes := range in.GetAttributes() {
		count := 1
		for key, value := range attributes.GetValues() {
			message := map[string]interface{}{
				"type":     reflect.ValueOf(value).Kind(),
				"key":      key,
				"value":    value,
				"position": count,
			}
			log.Printf("%v %v = %d", reflect.TypeOf(value), key, count)
			count += 1
			list = append(list, message)
		}
		jsonBytes, _ := protojson.Marshal(&userpb.Payload_Attributes{Values: attributes.GetValues()})
		fmt.Println(string(jsonBytes))
		fmt.Println(list)
	}

	return &userpb.Payload{Attributes: in.GetAttributes()}, nil
}

func (as *userServer) PostPayload(ctx context.Context, in *userpb.AnyPayload) (*userpb.Result, error) {

	rawDecodedText, err := base64.StdEncoding.DecodeString(in.GetBody())
	if err != nil {
		panic(err)
	}
	var result = map[string]interface{}{}
	rawError := json.Unmarshal(rawDecodedText, &result)

	if rawError != nil {
		panic(rawError)
	}
	log.Println(Prettify(result))

	mapString := make(map[string]string)
	for key, value := range result {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)
		mapString[strKey] = strValue
	}

	return &userpb.Result{Result: mapString}, nil
}

func Prettify(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func (as *userServer) CreateUser(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	user := models.User{Name: in.Name, Email: in.Email, PhoneNumber: in.PhoneNumber}

	database.NewRecord(user)
	database.Create(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) UpdateUserByName(ctx context.Context, in *userpb.User) (*userpb.User, error) {

	name := in.GetName()
	if name == "" {
		return &userpb.User{}, status.Error(codes.InvalidArgument, "User's name not specified")
	}
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
	if name == "" {
		return &userpb.Message{}, status.Error(codes.InvalidArgument, "User's name not specified")
	}
	var user models.User
	rowsAffected := database.Where("name =?", name).Find(&user).RowsAffected

	if rowsAffected == 0 {
		return &userpb.Message{}, status.Error(codes.NotFound, "Cannot find a User with this name!")
	}
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
	log.Println(colorPurple, "[userService] - [rpc GetUserByName] -> ",
		colorBlue, "Received person's name: ", in.GetName())
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	name := in.GetName()
	if name == "" {
		return &userpb.User{}, status.Error(codes.InvalidArgument, "User's name not specified")
	}
	var user models.User
	database.Where(&models.User{Name: name}).Find(&user)

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) GetUserById(ctx context.Context, in *userpb.User) (*userpb.User, error) {
	id := in.GetId()
	var user models.User
	rowsAffected := database.Where("id = ?", id).Find(&user).RowsAffected

	if rowsAffected == 0 {
		return &userpb.User{}, status.Error(codes.NotFound, "Cannot find a User with this id!")
	}

	return &userpb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func (as *userServer) FindUserFromGetUserByIdRPC(ctx context.Context, in *userpb.User) (*userpb.User, error) {

	user, _ := as.GetUserById(ctx, in)

	return &userpb.User{Id: user.GetId(), Name: user.GetName(), Email: user.GetEmail(),
		PhoneNumber: user.GetPhoneNumber()}, nil
}

func (as *userServer) FindUserFromGetUserByNameRPC(ctx context.Context, in *userpb.User) (*userpb.User, error) {

	user, _ := as.GetUserByName(ctx, in)

	return &userpb.User{Id: user.GetId(), Name: user.GetName(), Email: user.GetEmail(),
		PhoneNumber: user.GetPhoneNumber()}, nil
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

func (fu *creditCardServer) GetCreditCardByUserName(ctx context.Context,
	in *userpb.CreditCard) (*userpb.CreditCard, error) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	name := in.GetName()
	if name == "" {
		return &userpb.CreditCard{}, status.Error(codes.InvalidArgument, "User's name cannot be empty")
	}
	var creditCard models.CreditCards
	rowsAffected := database.Where(&models.CreditCards{Name: name}).Find(&creditCard).RowsAffected

	if rowsAffected == 0 {
		return &userpb.CreditCard{}, status.Error(codes.NotFound, "Cannot find a credit card with this user name!")
	}

	log.Println(colorPurple, "[creditCardService] - [rpc GetCreditCardByUserName] -> ", colorGreen,
		"Now sending the response (credit card of selected user) to client side...")

	return &userpb.CreditCard{Id: uint32(creditCard.ID), Name: creditCard.Name, Email: creditCard.Email,
		PhoneNumber: creditCard.PhoneNumber, Address: creditCard.Address, Country: creditCard.Country,
		City: creditCard.City, Zip: creditCard.Zip, Cvv: creditCard.CVV,
		CreatedAt: timestamppb.New(creditCard.CreatedAt)}, nil
}

func (fu *creditCardServer) CreateCreditCardApplication(ctx context.Context,
	in *userpb.CreditCardApplication) (*userpb.CreditCardApplication, error) {

	creditCardApplication := models.CreditCardApplication{
		FirstName:            in.GetFirstName(),
		LastName:             in.GetLastName(),
		DateOfBirth:          in.GetDateOfBirth().AsTime(),
		PhoneNumber:          in.GetPhoneNumber(),
		SocialSecurityNumber: in.GetSocialSecurityNumber(),
		EmploymentType:       in.GetEmploymentType(),
		Occupation:           in.GetOccupation(),
		MonthlyIncome:        float64(in.GetMonthlyIncome()),
		YearsEmployed:        int(in.GetYearsEmployed()),
		StreetAddress:        in.GetStreetAddress(),
		YearsAtAddress:       int(in.GetYearsAtAddress()),
		City:                 in.GetCity(),
		State:                in.GetState(),
		Zip:                  in.GetZip(),
		Country:              in.GetCountry(),
		Ownership:            in.GetOwnership(),
		MonthlyPayment:       float64(in.GetMonthlyPayment()),
		CardName:             in.GetCardName(),
		CardType:             in.GetCardType(),
		Branch:               in.GetBranch(),
		CardBranding:         in.GetCardBranding(),
	}

	database.NewRecord(creditCardApplication)
	database.Create(&creditCardApplication)

	return &userpb.CreditCardApplication{
		Id:                   uint32(creditCardApplication.ID),
		FirstName:            creditCardApplication.FirstName,
		LastName:             creditCardApplication.LastName,
		DateOfBirth:          timestamppb.New(creditCardApplication.DateOfBirth),
		PhoneNumber:          creditCardApplication.PhoneNumber,
		SocialSecurityNumber: creditCardApplication.SocialSecurityNumber,
		EmploymentType:       creditCardApplication.EmploymentType,
		Occupation:           creditCardApplication.Occupation,
		MonthlyIncome:        float32(creditCardApplication.MonthlyIncome),
		YearsEmployed:        int32(creditCardApplication.YearsEmployed),
		StreetAddress:        creditCardApplication.StreetAddress,
		YearsAtAddress:       int32(creditCardApplication.YearsAtAddress),
		City:                 creditCardApplication.City,
		State:                creditCardApplication.State,
		Zip:                  creditCardApplication.Zip,
		Country:              creditCardApplication.Country,
		Ownership:            creditCardApplication.Ownership,
		MonthlyPayment:       float32(creditCardApplication.MonthlyPayment),
		CardName:             creditCardApplication.CardName,
		CardType:             creditCardApplication.CardType,
		Branch:               creditCardApplication.Branch,
		CardBranding:         creditCardApplication.CardBranding,
		CreatedAt:            timestamppb.New(creditCardApplication.CreatedAt),
		UpdatedAt:            timestamppb.New(creditCardApplication.UpdatedAt),
		DeletedAt:            timestamppb.New(creditCardApplication.DeletedAt),
	}, nil
}

func (fu *creditCardServer) GetCreditCardApplicationByName(ctx context.Context,
	in *userpb.CreditCardApplication) (*userpb.CreditCardApplication, error) {

	firstName := in.GetFirstName()
	var creditCardApplication models.CreditCardApplication
	database.Unscoped().Where(&models.CreditCardApplication{FirstName: firstName}).
		Order("created_at desc").First(&creditCardApplication)

	return &userpb.CreditCardApplication{
		Id:                   uint32(creditCardApplication.ID),
		FirstName:            creditCardApplication.FirstName,
		LastName:             creditCardApplication.LastName,
		DateOfBirth:          timestamppb.New(creditCardApplication.DateOfBirth),
		PhoneNumber:          creditCardApplication.PhoneNumber,
		SocialSecurityNumber: creditCardApplication.SocialSecurityNumber,
		EmploymentType:       creditCardApplication.EmploymentType,
		Occupation:           creditCardApplication.Occupation,
		MonthlyIncome:        float32(creditCardApplication.MonthlyIncome),
		YearsEmployed:        int32(creditCardApplication.YearsEmployed),
		StreetAddress:        creditCardApplication.StreetAddress,
		YearsAtAddress:       int32(creditCardApplication.YearsAtAddress),
		City:                 creditCardApplication.City,
		State:                creditCardApplication.State,
		Zip:                  creditCardApplication.Zip,
		Country:              creditCardApplication.Country,
		Ownership:            creditCardApplication.Ownership,
		MonthlyPayment:       float32(creditCardApplication.MonthlyPayment),
		CardName:             creditCardApplication.CardName,
		CardType:             creditCardApplication.CardType,
		Branch:               creditCardApplication.Branch,
		CardBranding:         creditCardApplication.CardBranding,
		CreatedAt:            timestamppb.New(creditCardApplication.CreatedAt),
		UpdatedAt:            timestamppb.New(creditCardApplication.UpdatedAt),
		DeletedAt:            timestamppb.New(creditCardApplication.DeletedAt),
	}, nil
}

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

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the User service to the server
	userpb.RegisterUserServiceServer(s, &userServer{})
	userpb.RegisterCreditCardServiceServer(s, &creditCardServer{})

	// Serve gRPC server
	log.Printf("Serving gRPC on %s:%s", os.Getenv("SERVER_HOST"), os.Getenv("GRPC_SERVER_PORT"))
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	maxMsgSize := 1024 * 1024 * 20
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("GRPC_SERVER_PORT")),
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
		Addr:    fmt.Sprintf(":%s", os.Getenv("GRPC_GATEWAY_SERVER_PORT")),
		Handler: cors(gwmux),
	}

	log.Printf("Serving gRPC-Gateway on %s:%s", os.Getenv("SERVER_HOST"),
		os.Getenv("GRPC_GATEWAY_SERVER_PORT"))
	log.Fatalln(gwServer.ListenAndServe())
}
