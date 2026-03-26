package dto

import adminpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/admin"

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

func SignInRequestFromProto(req *adminpb.LoginRequest) SignInRequest {
	return SignInRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (r *SignInResponse) ToProto() *adminpb.LoginResponse {
	return &adminpb.LoginResponse{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
	}
}
