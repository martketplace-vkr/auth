package client

import (
	"context"
	"fmt"
	"time"

	"github.com/martketplace-vkr/auth/internal/service/client/dto"

	"github.com/redis/go-redis/v9"
)

const (
	RefreshKeyFmt = "refresh_token:%s"
)

type cache struct {
	cfg Config
	rd  *redis.Client
}

func New(rd *redis.Client) *cache {
	return &cache{
		rd: rd,
	}
}

func (c *cache) SaveRefreshToken(ctx context.Context, args dto.SaveRefreshTokenArgs) (err error) {
	return c.rd.Set(
		ctx,
		fmt.Sprintf(RefreshKeyFmt, args.RefreshHash),
		args.UserID,
		time.Duration(c.cfg.RerfreshTokenTtlSeconds*int64(time.Second)),
	).Err()
}

func (s *cache) GetUserByRefreshHash(
	ctx context.Context,
	tokenHash string,
) (int64, error) {
	val, err := s.rd.Get(ctx, fmt.Sprintf(RefreshKeyFmt, tokenHash)).Int64()
	if err != nil {
		return 0, err
	}

	return val, nil
}

func (s *cache) DeleteRefreshToken(
	ctx context.Context,
	tokenHash string,
) error {
	return s.rd.Del(ctx, fmt.Sprintf(RefreshKeyFmt, tokenHash)).Err()
}
