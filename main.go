package main

import (
	"log"
	"net"

	app "github.com/isaqueveraslabs/email-microservice/application/email"
	config "github.com/isaqueveraslabs/email-microservice/configuration"
	inter "github.com/isaqueveraslabs/email-microservice/interfaces/email"
	gogrpc "google.golang.org/grpc"
)

func main() {
	// Loading config of system
	config.Load()

	// Listen on port
	listen, err := net.Listen("tcp", config.Get().Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	// Creating new server
	server := gogrpc.NewServer()

	// Email registration server
	app.RegisterEMailServer(server, &inter.Server{})

	// Message of success
	log.Println(config.Get().Name, "is running in port", config.Get().Address)

	// Initializing server
	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err.Error())
	}
}
