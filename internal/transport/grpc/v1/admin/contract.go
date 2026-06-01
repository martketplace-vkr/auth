package admin

import (
	"context"

	"github.com/martketplace-vkr/auth/domain"
	"github.com/martketplace-vkr/auth/internal/service/admin/dto"
)

type (
	service interface {
		SignIn(ctx context.Context, req dto.SignInRequest) (resp dto.SignInResponse, err error)
		SignUp(ctx context.Context, req dto.SignUpRequest) (resp dto.SignUpResponse, err error)
		SignOut(ctx context.Context, req dto.SignOutRequest) (err error)
		ValidateToken(
			ctx context.Context,
			accessToken string,
		) (int64, string, error)
		Refresh(
			ctx context.Context,
			refreshToken string,
		) (resp dto.RefreshResponse, err error)
		CreateInviteToken(
			ctx context.Context,
			req dto.CreateInviteTokenRequest,
		) (resp dto.CreateInviteTokenResponse, err error)
		ListVendors(ctx context.Context) ([]domain.User, error)
		ListClients(ctx context.Context, query string, userStatus string, limit uint32, offset uint32) ([]domain.Client, uint64, error)
		GetClient(ctx context.Context, clientID int64) (domain.Client, error)
		UpdateClientStatus(ctx context.Context, clientID int64, userStatus string, reason string, adminID int64) (domain.Client, error)
	}
)
