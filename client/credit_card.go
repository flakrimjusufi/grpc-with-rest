package client

import (
	"context"
	"log"

	db "github.com/flakrimjusufi/grpc-with-rest/database"
	"github.com/flakrimjusufi/grpc-with-rest/models"
	creditpb "github.com/flakrimjusufi/grpc-with-rest/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	colorGreen  = "\033[32m"
	colorPurple = "\033[35m"
)

var database = db.Connect().Debug()

// CreditCardServer - the grpc server of credit cards
type CreditCardServer struct {
	creditpb.UnimplementedCreditCardServiceServer
}

// CreditCards - the service that gets a list of credit cards by interacting with models.CreditCards and returns a creditpb.ListCreditCards as a response
func (cs *CreditCardServer) CreditCards(ctx context.Context, in *creditpb.CreditCard) (*creditpb.ListCreditCards, error) {

	var list []*creditpb.CreditCard
	var creditCards []*models.CreditCards
	database.Order("created_at desc").Find(&creditCards)

	for _, card := range creditCards {
		list = append(list, &creditpb.CreditCard{
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

	return &creditpb.ListCreditCards{
		CreditCards: list,
	}, nil
}

// GetCreditCardByUserName - the service that gets the credit card by userName with models.CreditCards and returns a creditpb.CreditCard as a response
func (cs *CreditCardServer) GetCreditCardByUserName(ctx context.Context,
	in *creditpb.CreditCard) (*creditpb.CreditCard, error) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	name := in.GetName()
	if name == "" {
		return &creditpb.CreditCard{}, status.Error(codes.InvalidArgument, "User's name cannot be empty")
	}
	var creditCard models.CreditCards
	rowsAffected := database.Where(&models.CreditCards{Name: name}).Find(&creditCard).RowsAffected

	if rowsAffected == 0 {
		return &creditpb.CreditCard{}, status.Error(codes.NotFound, "Cannot find a credit card with this user name!")
	}

	log.Println(colorPurple, "[creditCardService] - [rpc GetCreditCardByUserName] -> ", colorGreen,
		"Now sending the response (credit card of selected user) to client side...")

	return &creditpb.CreditCard{Id: uint32(creditCard.ID), Name: creditCard.Name, Email: creditCard.Email,
		PhoneNumber: creditCard.PhoneNumber, Address: creditCard.Address, Country: creditCard.Country,
		City: creditCard.City, Zip: creditCard.Zip, Cvv: creditCard.CVV,
		CreatedAt: timestamppb.New(creditCard.CreatedAt)}, nil
}

// CreateCreditCardApplication - the service that creates the credit card application with models.CreditCardApplication and returns a creditpb.CreditCardApplication as a response
func (cs *CreditCardServer) CreateCreditCardApplication(ctx context.Context,
	in *creditpb.CreditCardApplication) (*creditpb.CreditCardApplication, error) {

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

	return &creditpb.CreditCardApplication{
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

// GetCreditCardApplicationByName - the service that gets the credit card from models.CreditCardApplication and returns a creditpb.CreditCardApplication as a response
func (cs *CreditCardServer) GetCreditCardApplicationByName(ctx context.Context,
	in *creditpb.CreditCardApplication) (*creditpb.CreditCardApplication, error) {

	firstName := in.GetFirstName()
	var creditCardApplication models.CreditCardApplication
	database.Unscoped().Where(&models.CreditCardApplication{FirstName: firstName}).
		Order("created_at desc").First(&creditCardApplication)

	return &creditpb.CreditCardApplication{
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
