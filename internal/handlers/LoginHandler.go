package handlers

import (
	"context"

	"github.com/DogGoOrg/doggo-account/internal/helpers"
	"github.com/DogGoOrg/doggo-account/proto/proto_services/Account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// type tokenPayload struct {
// 	Email, Password string
// }

// var accessSecret = []byte("SecretYouShouldHide")

// var refreshSecret = []byte("SecretYouShoulRefdHide")

// func genTokenPair(payload tokenPayload) ([2]string, error) {
// 	arr := [2]string{}

// 	accessToken := jwt.New(jwt.SigningMethodPS256.SigningMethodRSA)
// 	// refreshToken := jwt.New(jwt.SigningMethodPS256.SigningMethodRSA)

// 	accessClaims := accessToken.Claims.(jwt.MapClaims)
// 	accessClaims["exp"] = time.Now().Add(10 * time.Minute)

// 	return arr, nil
// }

func LoginHandler(ctx context.Context, in *Account.LoginRequest) (*Account.LoginResponse, error) {
	email, password := in.Email, in.Password

	if err := helpers.CheckForNullValues[string](email, password); err != nil {
		//TODO handle validation error
		return nil, status.Error(codes.InvalidArgument, "LOL")
	}

	return &Account.LoginResponse{
		AccessToken:  "",
		RefreshToken: "",
		Id:           "1",
		Email:        "nik@gmail.com",
	}, nil
}
