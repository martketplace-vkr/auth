package client

import (
	"context"

	"github.com/martketplace-vkr/auth/internal/service/client/dto"
	"github.com/martketplace-vkr/auth/pkg/api/grpc/v1/client"
)

type Handler struct {
	service service
	client.AuthClientServiceServer
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (s *Handler) Register(ctx context.Context, req *client.RegisterRequest) (resp *client.RegisterResponse, err error) {
	response, err := s.service.SignUp(ctx, dto.SignUpRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (s *Handler) Login(ctx context.Context, req *client.LoginRequest) (resp *client.LoginResponse, err error) {
	return resp, nil
}

func (s *Handler) ValidateToken(ctx context.Context, req *client.ValidateTokenRequest) (resp *client.ValidateTokenResponse, err error) {
	return resp, nil
}
