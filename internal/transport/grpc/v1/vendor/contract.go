package vendor

import (
	"context"

	vendordto "github.com/martketplace-vkr/auth/internal/service/vendor_dto"
)

type (
	service interface {
		SignIn(ctx context.Context, req vendordto.SignInRequest) (resp vendordto.SignInResponse, err error)
		SignUp(ctx context.Context, req vendordto.SignUpRequest) (resp vendordto.SignUpResponse, err error)
		SignOut(ctx context.Context, req vendordto.SignOutRequest) (err error)
		ValidateToken(
			ctx context.Context,
			accessToken string,
		) (int64, error)
		Refresh(
			ctx context.Context,
			refreshToken string,
		) (resp vendordto.RefreshResponse, err error)
	}
)
