package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/DogGoOrg/doggo-account/proto/proto_services/Account"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	Account.UnimplementedAccountServer
}

// ping method implementation
func (receiver *server) Ping(ctx context.Context, in *Account.PingRequest) (*Account.PingResponse, error) {
	fmt.Println("OK")
	return &Account.PingResponse{
		Status: "OK",
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	Account.RegisterAccountServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
