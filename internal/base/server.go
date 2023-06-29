package base

import (
	"context"

	"github.com/DogGoOrg/doggo-account/proto/proto_services/Account"
	"gorm.io/gorm"
)

type Server struct {
	Account.UnimplementedAccountServer
	Database *gorm.DB
}

func (r *Server) Ping(ctx context.Context, in *Account.PingRequest) (*Account.PingResponse, error) {
	return &Account.PingResponse{
		Status: "OK",
	}, nil
}

func (r *Server) GetAccountById(ctx context.Context, in *Account.GetAccountRequest) (*Account.GetAccountResponse, error) {
	return &Account.GetAccountResponse{
		Id:   "122",
		Info: "No info",
	}, nil
}

func (r *Server) Login(ctx context.Context, in *Account.LoginRequest) (*Account.LoginResponse, error) {
	return &Account.LoginResponse{
		AccessToken:  "",
		RefreshToken: "",
		Id:           "1",
		Email:        "nik@gmail.com",
	}, nil
}

func (r *Server) Logout(ctx context.Context, in *Account.LogoutRequest) (*Account.LogoutResponse, error) {
	return &Account.LogoutResponse{
		Status: "OK",
	}, nil
}

func (r *Server) Refresh(ctx context.Context, in *Account.RefreshRequest) (*Account.RefreshResponse, error) {
	return &Account.RefreshResponse{
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}
