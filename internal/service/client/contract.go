package client

import (
	"context"

	"github.com/martketplace-vkr/auth/domain"
	"github.com/martketplace-vkr/auth/internal/service/client/dto"
)

type (
	txManager interface {
		Do(ctx context.Context, fn func(ctx context.Context) error) (err error)
	}
	repository interface {
		InsertUser(ctx context.Context, user *domain.User) error
		SelectUserByEmail(ctx context.Context, email string) (user *domain.User, err error)
	}
	cache interface {
		SaveRefreshToken(ctx context.Context, args dto.SaveRefreshTokenArgs) (err error)
		GetUserByRefreshHash(
			ctx context.Context,
			tokenHash string,
		) (int64, error)
		DeleteRefreshToken(
			ctx context.Context,
			tokenHash string,
		) error
	}
)
