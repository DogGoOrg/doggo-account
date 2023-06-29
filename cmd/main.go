package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/DogGoOrg/doggo-account/internal/db"
	"github.com/DogGoOrg/doggo-account/proto/proto_services/Account"
	"github.com/DogGoOrg/doggo-orm/models"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type server struct {
	Account.UnimplementedAccountServer
	Database *gorm.DB
}

// ping method implementation
func (receiver *server) Ping(ctx context.Context, in *Account.PingRequest) (*Account.PingResponse, error) {
	return &Account.PingResponse{
		Status: "OK",
	}, nil
}

func main() {
	appMode := os.Getenv("MODE")

	if appMode != "prod" {
		if err := godotenv.Load("dev.env"); err != nil {
			panic(err.Error())
		}
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serv := new(server)

	db := db.InitDB()

	serv.Database = db

	db.AutoMigrate(&models.Account{}, &models.Pet{})

	s := grpc.NewServer()

	Account.RegisterAccountServer(s, serv)
	log.Printf("server listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
