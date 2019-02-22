package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// CreateEmail generates response to a Ping request
func (s *Server) CreateEmail(ctx context.Context, in *Email) (*CreateEmailResponse, error) {
	log.Printf("Receive message %s, %s, %s", in.To, in.Subject, in.Message)
	return &CreateEmailResponse{
		Code:    201,
		Message: "Success created item",
	}, nil
}
