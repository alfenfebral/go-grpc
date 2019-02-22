package main

import (
	"log"

	"grpc-tst/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewEmailHandlerClient(conn)
	response, err := c.CreateEmail(context.Background(), &api.Email{To: "foo@foo.foo", Subject: "foo", Message: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %v, %s", response.Code, response.Message)
}
