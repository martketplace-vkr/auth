package seller

import (
	"context"

	"github.com/martketplace-vkr/auth/internal/service/seller/dto"
	vendorpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/vendor"
)

type Handler struct {
	service service
	vendorpb.AuthVendorServiceServer
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(ctx context.Context, req *vendorpb.RegisterRequest) (resp *vendorpb.RegisterResponse, err error) {
	response, err := h.service.SignUp(ctx, dto.SignUpRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) Login(ctx context.Context, req *vendorpb.LoginRequest) (resp *vendorpb.LoginResponse, err error) {
	response, err := h.service.SignIn(ctx, dto.SignInRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *vendorpb.ValidateTokenRequest) (resp *vendorpb.ValidateTokenResponse, err error) {
	vendorID, err := h.service.ValidateToken(ctx, req.Token)
	if err != nil {
		return resp, err
	}

	return &vendorpb.ValidateTokenResponse{
		VendorId: vendorID,
		Role:     "vendor",
	}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *vendorpb.RefreshTokenRequest) (resp *vendorpb.RefreshTokenResponse, err error) {
	response, err := h.service.Refresh(ctx, req.RefreshToken)
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) Logout(ctx context.Context, req *vendorpb.LogoutRequest) (resp *vendorpb.LogoutResponse, err error) {
	err = h.service.SignOut(ctx, dto.SignOutRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return resp, err
	}

	return &vendorpb.LogoutResponse{
		Success: true,
	}, nil
}
