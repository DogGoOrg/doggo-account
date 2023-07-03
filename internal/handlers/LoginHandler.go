package handlers

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/DogGoOrg/doggo-account/internal/helpers"
	"github.com/DogGoOrg/doggo-account/proto/proto_services/Account"
	"github.com/DogGoOrg/doggo-orm/models"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func genTokenPair(email string, id any) (*[2]string, error) {
	arr := [2]string{}

	accessClaims := &jwt.MapClaims{
		"email": email,
		"id":    id,
		"expIn": time.Now().Add(15 * time.Minute),
	}

	refreshClaims := &jwt.MapClaims{
		"expIn": time.Now().Add(time.Minute * 60 * 24 * 30), // 30 days
		"email": email,
		"id":    id,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessSecret, refreshSecret := os.Getenv("ACCESS_SECRET"), os.Getenv("REFRESH_SECRET")

	if access, err := accessToken.SignedString([]byte(accessSecret)); err == nil {
		arr[0] = access
	} else {
		return nil, err
	}

	if refreshToken, err := refreshToken.SignedString([]byte(refreshSecret)); err == nil {
		arr[1] = refreshToken
	} else {
		return nil, err
	}

	return &arr, nil
}

func getPasswordHash(password string) string {
	hasher := sha1.New()

	if _, err := hasher.Write([]byte(password)); err == nil {
		hash := hex.EncodeToString(hasher.Sum(nil))
		return hash
	}

	return ""

}

func LoginHandler(ctx context.Context, in *Account.LoginRequest, db *gorm.DB) (*Account.LoginResponse, error) {
	// log.Fatalln("AAAAÀ")
	//get response parameters
	email, password := in.Email, in.Password

	if err := helpers.CheckForNullValues[string](email, password); err != nil {
		//TODO handle validation error
		return nil, status.Error(codes.InvalidArgument, "LOL")
	}

	//TODO try get user from database
	wg := &sync.WaitGroup{}
	var account *models.Account = &models.Account{Email: email}

	wg.Add(1)
	go func(email string) {
		db.Take(&account)
		wg.Done()
	}(email)

	wg.Wait()

	if account == nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	//get hashed password
	hash := getPasswordHash(password)

	//compare hashed password with stored password
	if hash != account.Password {
		return nil, status.Error(codes.Unauthenticated, "Wrong password")
	}

	var tokenPair *[2]string

	//generate token pair
	if tokens, err := genTokenPair(email, account.ID); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		tokenPair = tokens
	}

	//return response object
	return &Account.LoginResponse{
		AccessToken:  tokenPair[0],
		RefreshToken: tokenPair[1],
		Id:           fmt.Sprintf("%v", account.ID),
		Email:        email,
	}, nil
}
