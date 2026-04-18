package admin

import (
	"context"

	"github.com/martketplace-vkr/auth/domain"
	clientdto "github.com/martketplace-vkr/auth/internal/service/client/dto"
)

type (
	txManager interface {
		Do(ctx context.Context, fn func(ctx context.Context) error) (err error)
	}
	repository interface {
		InsertUser(ctx context.Context, user *domain.User) error
		SelectUserByEmail(ctx context.Context, email string) (user *domain.User, err error)
		SelectUserByID(ctx context.Context, id int64) (user *domain.User, err error)
		HasActiveInviteToken(ctx context.Context, token string) (bool, error)
		UseInviteToken(ctx context.Context, token string, userID int64) error
		CreateInviteToken(ctx context.Context, createdBy int64, roleID int64, token string) error
	}
	cache interface {
		SaveRefreshToken(ctx context.Context, args clientdto.SaveRefreshTokenArgs) (err error)
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
