package handlers

import (
	"context"

	"github.com/DogGoOrg/doggo-account/proto/proto_services/Account"
)

func PingHandler(ctx context.Context, in *Account.PingRequest) (*Account.PingResponse, error) {
	return &Account.PingResponse{
		Status: "OK",
	}, nil
}
