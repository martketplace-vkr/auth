package dto

import "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/client"

type (
	SignInRequest struct {
		Email    string
		Password string
	}
	SignInResponse struct {
		AccessToken  string
		RefreshToken string
	}
)

func SignInRequestFromProto(req *client.LoginRequest) SignInRequest {
	return SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (r *SignInResponse) ToProto() *client.LoginResponse {
	return &client.LoginResponse{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	}
}

type (
	SaveRefreshTokenArgs struct {
		RefreshHash string
		UserID      int64
	}
)
