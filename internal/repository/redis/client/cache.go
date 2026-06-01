package client

import (
	"context"
	"fmt"
	"time"

	"github.com/martketplace-vkr/auth/internal/service/client/dto"

	"github.com/redis/go-redis/v9"
)

const (
	RefreshKeyFmt      = "refresh_token:%s"
	UserRefreshKeyFmt  = "refresh_tokens:user:%d"
	UserActivityKeyFmt = "activity_recorded:user:%d:%s"
)

type cache struct {
	cfg Config
	rd  *redis.Client
}

func (s *cache) MarkActivity(ctx context.Context, userID int64) (bool, error) {
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return false, err
	}
	key := fmt.Sprintf(UserActivityKeyFmt, userID, time.Now().In(location).Format("2006-01-02"))
	return s.rd.SetNX(ctx, key, "1", 48*time.Hour).Result()
}

func New(rd *redis.Client) *cache {
	cfg := Config{
		RerfreshTokenTtlSeconds: 60 * 60 * 24 * 30,
	}

	return &cache{
		cfg: cfg,
		rd:  rd,
	}
}

func (c *cache) SaveRefreshToken(ctx context.Context, args dto.SaveRefreshTokenArgs) (err error) {
	ttl := time.Duration(c.cfg.RerfreshTokenTtlSeconds * int64(time.Second))
	tokenKey := fmt.Sprintf(RefreshKeyFmt, args.RefreshHash)
	userKey := fmt.Sprintf(UserRefreshKeyFmt, args.UserID)
	pipe := c.rd.TxPipeline()
	pipe.Set(ctx, tokenKey, args.UserID, ttl)
	pipe.SAdd(ctx, userKey, args.RefreshHash)
	pipe.Expire(ctx, userKey, ttl)
	_, err = pipe.Exec(ctx)
	return err
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

func (s *cache) DeleteUserRefreshTokens(ctx context.Context, userID int64) error {
	userKey := fmt.Sprintf(UserRefreshKeyFmt, userID)
	hashes, err := s.rd.SMembers(ctx, userKey).Result()
	if err != nil {
		return err
	}

	keys := make([]string, 0, len(hashes)+1)
	keys = append(keys, userKey)
	for _, hash := range hashes {
		keys = append(keys, fmt.Sprintf(RefreshKeyFmt, hash))
	}

	return s.rd.Del(ctx, keys...).Err()
}
