package main

import (
	"log"

	pb "github.com/RichJames/shippy/user-service/proto/auth"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/mdns"
)

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automtically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service.  Optionally, include some options here.
	srv := micro.NewService(

		// This name maybe should match the package name given in the protobuf definition
		micro.Name("shippy.auth"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// publisher := micro.NewPublisher("user.created", srv.Client())


	// Register handler
	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
