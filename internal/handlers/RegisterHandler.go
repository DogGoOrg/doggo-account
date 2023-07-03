package handlers

import (
	"context"

	"github.com/DogGoOrg/doggo-account/internal/helpers"
	"github.com/DogGoOrg/doggo-account/proto/proto_services/Account"
	"github.com/DogGoOrg/doggo-orm/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func RegisterHandler(ctx context.Context, in *Account.RegisterRequest, db *gorm.DB) (*Account.RegisterResponse, error) {
	email, password := in.Email, in.Password

	if err := helpers.CheckForNullValues[string](email, password); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid email or password")
	}

	dbCh := make(chan int, 1)
	defer close(dbCh)

	var acc models.Account

	go func(email string, ch chan<- int) {
		res := db.Limit(1).Find(&acc, "email =?", email)
		ch <- int(res.RowsAffected)
	}(email, dbCh)

	res := <-dbCh

	if res > 0 {
		return nil, status.Error(codes.AlreadyExists, "email already exists")
	}

	passHash := helpers.GetPasswordHash(password)

	newAccount := &models.Account{Email: email, Password: passHash}
	db.Create(newAccount)

	return &Account.RegisterResponse{Status: "Success"}, nil
}
