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

func (h *Handler) Register(ctx context.Context, req *client.RegisterRequest) (resp *client.RegisterResponse, err error) {
	response, err := h.service.SignUp(ctx, dto.SignUpRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) Login(ctx context.Context, req *client.LoginRequest) (resp *client.LoginResponse, err error) {
	response, err := h.service.SignIn(ctx, dto.SignInRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *client.ValidateTokenRequest) (resp *client.ValidateTokenResponse, err error) {
	userID, err := h.service.ValidateToken(ctx, req.Token)
	if err != nil {
		return resp, err
	}

	return &client.ValidateTokenResponse{
		UserId: userID,
	}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *client.RefreshTokenRequest) (resp *client.RefreshTokenResponse, err error) {
	response, err := h.service.Refresh(ctx, req.RefreshToken)
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) Logout(ctx context.Context, req *client.LogoutRequest) (resp *client.LogoutResponse, err error) {
	err = h.service.SignOut(ctx, dto.SignOutRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return resp, err
	}

	return &client.LogoutResponse{
		Success: true,
	}, nil
}
