package client

import "github.com/martketplace-vkr/pkg/outbox"

type service struct {
	cfg        Config
	txManager  txManager
	repository repository
	cache      cache
	outbox     outbox.Outbox
}

func New(
	cfg Config,
	txManager txManager,
	repository repository,
	cache cache,
	outbox outbox.Outbox,
) *service {
	return &service{
		cfg:        cfg,
		txManager:  txManager,
		repository: repository,
		cache:      cache,
		outbox:     outbox,
	}
}
