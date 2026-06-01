package admin

import (
	"context"
	"time"

	"github.com/martketplace-vkr/auth/domain"
	admindto "github.com/martketplace-vkr/auth/internal/service/admin/dto"
	adminpb "github.com/martketplace-vkr/auth/pkg/api/grpc/v1/admin"
)

type Handler struct {
	service service
	adminpb.AuthAdminServiceServer
}

func (h *Handler) ListClients(ctx context.Context, req *adminpb.ListClientsRequest) (*adminpb.ListClientsResponse, error) {
	clients, total, err := h.service.ListClients(ctx, req.GetQuery(), req.GetStatus(), req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, err
	}
	resp := &adminpb.ListClientsResponse{Total: total, Clients: make([]*adminpb.Client, 0, len(clients))}
	for _, client := range clients {
		resp.Clients = append(resp.Clients, clientToProto(client))
	}
	return resp, nil
}

func (h *Handler) GetClient(ctx context.Context, req *adminpb.GetClientRequest) (*adminpb.Client, error) {
	client, err := h.service.GetClient(ctx, req.GetClientId())
	if err != nil {
		return nil, err
	}
	return clientToProto(client), nil
}

func (h *Handler) UpdateClientStatus(ctx context.Context, req *adminpb.UpdateClientStatusRequest) (*adminpb.Client, error) {
	client, err := h.service.UpdateClientStatus(ctx, req.GetClientId(), req.GetStatus(), req.GetReason(), req.GetAdminId())
	if err != nil {
		return nil, err
	}
	return clientToProto(client), nil
}

func clientToProto(client domain.Client) *adminpb.Client {
	resp := &adminpb.Client{
		ClientId:      client.ID,
		Email:         client.Email,
		EmailVerified: client.EmailVerified,
		Status:        client.Status,
		CreatedAt:     formatTime(client.CreatedAt),
	}
	if client.StatusReason != nil {
		resp.StatusReason = *client.StatusReason
	}
	if client.UpdatedAt != nil {
		resp.UpdatedAt = formatTime(*client.UpdatedAt)
	}
	if client.LastActivityAt != nil {
		resp.LastActivityAt = formatTime(*client.LastActivityAt)
	}
	for _, event := range client.ModerationEvents {
		resp.ModerationEvents = append(resp.ModerationEvents, &adminpb.ClientModerationEvent{
			Id: event.ID, AdminId: event.AdminID, OldStatus: event.OldStatus,
			NewStatus: event.NewStatus, Reason: event.Reason, CreatedAt: formatTime(event.CreatedAt),
		})
	}
	return resp
}

func formatTime(value time.Time) string {
	if value.IsZero() {
		return ""
	}
	return value.UTC().Format(time.RFC3339)
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
