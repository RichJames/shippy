// shippy/user-cli/cli.go
package main

import (
	"log"
	"os"

	pb "github.com/RichJames/shippy/user-service/proto/user"
	micro "github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

func main() {

	srv := micro.NewService(

		micro.Name("user-cli"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	client := pb.NewUserServiceClient("user.srv", microclient.DefaultClient)

	name := "Rich James"
	email := "richjamespub1@gmail.com"
	password := "test123"
	company := "Black Dog Company LLC"


	r, err := client.Create(context.TODO(), &pb.User{
		Name: 		name,
		Email: 		email,
		Password: 	password,
		Company: 	company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:		email,
		Password:	password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// Let's just exit because
	os.Exit(0)
}
