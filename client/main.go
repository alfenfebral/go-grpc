package main

import (
	"grpc-tst/api"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Authentication holds the login/password
type Authentication struct {
	Login    string
	Password string
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}
func main() {
	var conn *grpc.ClientConn
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}
	// Setup the login/pass
	auth := Authentication{
		Login:    "john",
		Password: "doe",
	}
	// Initiate a connection with the server
	conn, err = grpc.Dial("localhost:7777", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewEmailClient(conn)
	response, err := c.CreateEmail(context.Background(), &api.CreateEmailRequest{To: "foo@foo.foo", Subject: "foo", Message: "foo"})
	if err != nil {
		log.Fatalf("Error when calling CreateEmail: %s", err)
	}
	log.Printf("Response from server: %v, %s", response.Code, response.Message)
}
