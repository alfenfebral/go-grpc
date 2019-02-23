package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// CreateEmail - create email
func (s *Server) CreateEmail(ctx context.Context, in *CreateEmailRequest) (*CreateEmailResponse, error) {
	log.Printf("Receive message %s, %s, %s", in.To, in.Subject, in.Message)
	return &CreateEmailResponse{
		Code:    201,
		Message: "Success created item",
	}, nil
}

// GetEmail - get single email
func (s *Server) GetEmail(ctx context.Context, in *GetEmailRequest) (*GetEmailResponse, error) {
	log.Printf("Receive message %v", in.Id)
	return &GetEmailResponse{
		Id: in.Id,
	}, nil
}
