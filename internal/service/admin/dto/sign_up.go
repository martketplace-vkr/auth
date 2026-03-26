package dto

import (
	"github.com/martketplace-vkr/auth/domain"
	adminpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/admin"
)

type (
	SignUpRequest struct {
		Email       string
		Password    string
		InviteToken string
	}
	SignUpResponse struct {
		UserID int64
	}
)

func SignUpRequestFromProto(req *adminpb.RegisterRequest) SignUpRequest {
	return SignUpRequest{
		Email:       req.Email,
		Password:    req.Password,
		InviteToken: req.InviteToken,
	}
}

func (r *SignUpRequest) ToDomain() *domain.User {
	return &domain.User{
		Email:        r.Email,
		PasswordHash: r.Password,
	}
}

func (r *SignUpResponse) ToProto() *adminpb.RegisterResponse {
	return &adminpb.RegisterResponse{
		UserId: r.UserID,
	}
}
