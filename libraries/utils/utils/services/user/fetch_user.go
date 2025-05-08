package user

import (
	"context"
	"net/http"
)

type FetchUserResponse struct {
	user UserSerializer
}

type FetchUserRequest struct {
	Auth0ID      string
	PhoneNumber  string
	VerifyNumber string
}

func FetchUser(context context.Context, params FetchUserRequest) (FetchUserResponse, error) {
	return FetchUserResponse{}, nil
}

func FetchUserParams(req *http.Request) (FetchUserRequest, error) {
}
