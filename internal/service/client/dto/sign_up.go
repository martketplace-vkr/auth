package dto

import (
	"github.com/martketplace-vkr/auth/domain"
	"github.com/martketplace-vkr/auth/pkg/api/grpc/v1/client"
)

type (
	SignUpRequest struct {
		Email    string
		Password string
	}
	SignUpResponse struct {
		UserId int64
	}
)

func SignUpRequestFromProto(req *client.RegisterRequest) SignUpRequest {
	return SignUpRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (r *SignUpRequest) ToDomain() *domain.User {
	return &domain.User{
		Email:        r.Email,
		PasswordHash: r.Password,
	}
}

func (r *SignUpResponse) ToProto() *client.RegisterResponse {
	return &client.RegisterResponse{
		UserId: r.UserId,
	}
}
