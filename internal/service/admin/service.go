package admin

import (
	"context"

	"github.com/martketplace-vkr/auth/domain"
)

type service struct {
	cfg        Config
	txManager  txManager
	repository repository
	cache      cache
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
