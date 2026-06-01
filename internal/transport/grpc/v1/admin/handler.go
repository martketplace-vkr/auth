package admin

import (
	"context"

	admindto "github.com/martketplace-vkr/auth/internal/service/admin/dto"
	adminpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/admin"
)

type Handler struct {
	service service
	adminpb.AuthAdminServiceServer
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(ctx context.Context, req *adminpb.RegisterRequest) (resp *adminpb.RegisterResponse, err error) {
	response, err := h.service.SignUp(ctx, admindto.SignUpRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) Login(ctx context.Context, req *adminpb.LoginRequest) (resp *adminpb.LoginResponse, err error) {
	response, err := h.service.SignIn(ctx, admindto.SignInRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *adminpb.ValidateTokenRequest) (resp *adminpb.ValidateTokenResponse, err error) {
	userID, login, err := h.service.ValidateToken(ctx, req.Token)
	if err != nil {
		return resp, err
	}

	return &adminpb.ValidateTokenResponse{
		UserId: userID,
		Role:   "admin",
		Login:  login,
	}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *adminpb.RefreshTokenRequest) (resp *adminpb.RefreshTokenResponse, err error) {
	response, err := h.service.Refresh(ctx, req.RefreshToken)
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) Logout(ctx context.Context, req *adminpb.LogoutRequest) (resp *adminpb.LogoutResponse, err error) {
	err = h.service.SignOut(ctx, admindto.SignOutRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return resp, err
	}

	return &adminpb.LogoutResponse{
		Success: true,
	}, nil
}

func (h *Handler) CreateInviteToken(ctx context.Context, req *adminpb.CreateInviteTokenRequest) (resp *adminpb.CreateInviteTokenResponse, err error) {
	response, err := h.service.CreateInviteToken(ctx, admindto.CreateInviteTokenRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	return response.ToProto(), nil
}

func (h *Handler) ListVendors(ctx context.Context, _ *adminpb.ListVendorsRequest) (*adminpb.ListVendorsResponse, error) {
	vendors, err := h.service.ListVendors(ctx)
	if err != nil {
		return nil, err
	}

	resp := &adminpb.ListVendorsResponse{
		Vendors: make([]*adminpb.Vendor, 0, len(vendors)),
	}
	for _, vendor := range vendors {
		resp.Vendors = append(resp.Vendors, &adminpb.Vendor{
			VendorId: vendor.ID,
			Email:    vendor.Email,
		})
	}

	return resp, nil
}
