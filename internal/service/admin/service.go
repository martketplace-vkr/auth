package admin

import (
	"context"
	"fmt"
	"strings"

	"github.com/martketplace-vkr/auth/domain"
)

type service struct {
	cfg        Config
	txManager  txManager
	repository repository
	cache      cache
}

func (s *service) ListClients(ctx context.Context, query string, userStatus string, limit uint32, offset uint32) ([]domain.Client, uint64, error) {
	if userStatus != "" && userStatus != "active" && userStatus != "blocked" {
		return nil, 0, fmt.Errorf("unsupported client status")
	}
	return s.repository.ListClients(ctx, query, userStatus, limit, offset)
}

func (s *service) GetClient(ctx context.Context, clientID int64) (domain.Client, error) {
	if clientID <= 0 {
		return domain.Client{}, fmt.Errorf("client_id must be positive")
	}
	return s.repository.GetClient(ctx, clientID)
}

func (s *service) UpdateClientStatus(ctx context.Context, clientID int64, userStatus string, reason string, adminID int64) (client domain.Client, err error) {
	reason = strings.TrimSpace(reason)
	if clientID <= 0 || adminID <= 0 {
		return client, fmt.Errorf("client_id and admin_id must be positive")
	}
	if userStatus != "active" && userStatus != "blocked" {
		return client, fmt.Errorf("unsupported client status")
	}
	if reason == "" {
		return client, fmt.Errorf("moderation reason is required")
	}
	err = s.txManager.Do(ctx, func(ctx context.Context) error {
		var updateErr error
		client, updateErr = s.repository.UpdateClientStatus(ctx, clientID, userStatus, reason, adminID)
		return updateErr
	})
	if err != nil {
		return client, err
	}
	if err = s.cache.DeleteUserRefreshTokens(ctx, clientID); err != nil {
		return client, err
	}
	return client, nil
}

func New(
	cfg Config,
	txManager txManager,
	repository repository,
	cache cache,
) *service {
	return &service{
		cfg:        cfg,
		txManager:  txManager,
		repository: repository,
		cache:      cache,
	}
}

func (s *service) ListVendors(ctx context.Context) ([]domain.User, error) {
	return s.repository.ListVendors(ctx)
}
