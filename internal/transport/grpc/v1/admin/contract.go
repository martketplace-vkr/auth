package admin

import (
	"context"

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
	}
)
