package main

import (
	"context"
	"fmt"
	userpb "github.com/flakrimjusufi/grpc-with-rest/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:8080"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorWhite  = "\033[37m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorYellow = "\033[33m"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := userpb.NewCreditCardServiceClient(conn)
	u := userpb.NewUserServiceClient(conn)

	//user client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	personName := "Flakrim"

	log.Println(colorRed, "Starting the server on client side...")
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	log.Printf("Sending person name [%v] to [UserService] | communicating with rpc call [GetUserByName] |", personName)
	x, err := u.GetUserByName(ctx, &userpb.User{
		Name: personName,
	})

	if err != nil {
		log.Fatalf("could not get person data: %v", err)
	}

	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	log.Println(colorGreen, "Successfully received person data from Server Side [UserService] | rpc call [GetUserByName] |:")
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	log.Println(colorRed, "Response from the server side [userService] - [rpc GetUserByName] : ")
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	fmt.Println(colorCyan, x)

	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	log.Println(
		colorWhite,
		"Now sending the retrieved data to [CreditCardService] in order interact with rpc call "+
			"[GetCreditCardByUserName] and to receive person's credit card information")

	//credit card client
	ctx2, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCreditCardByUserName(ctx2, &userpb.CreditCard{
		Name: x.GetName(),
	})
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	if err != nil {
		log.Fatalf("could not get any data: %v", err)
	}
	log.Println(
		colorPurple,
		"Successfully received Credit Card information of selected user from [CreditCardService] "+
			"by interacting with rpc call [GetUserByName] which belongs to [userService]",
	)

	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	log.Println(colorRed, "Response from the server side [creditCardService] - [rpc GetCreditCardByUserName] : ")
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
	fmt.Println(colorCyan, r)
	fmt.Println(colorYellow, "__________________________________________________________________________________"+
		"_______________________________________________________________________________________________________")
}
