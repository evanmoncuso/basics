package basics

import (
	context "context"
	"log"
)

type Server struct{}

func (s *Server) GetWelcome(ctx context.Context) (*Greeting, error) {
	log.Printf("The server got a request for a greeting!", ctx)

	return &Greeting{Body: "Hello from the server"}, nil
}
