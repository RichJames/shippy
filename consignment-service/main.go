// shippy/consignment-service/main.go
package main

import (
	// Import the generated protobuf code
	"fmt"
	"log"
	pb "github.com/RichJames/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/RichJames/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"os"
)

const (
	defaultHost = "localhost:27017"
)


func main() {

	// Database host from the environment variables
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)

	// Mgo creates a 'master' session, we need to close that session
	// before the main function closes.
	defer session.Close()

	if err != nil {
		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	// Create a new service.  Optionally include some options here.
	srv := micro.NewService(

		// This name must match(??) the "package" name given in your protobuf definition
		micro.Name("consignment-server"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselServiceClient("vessel-service", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
