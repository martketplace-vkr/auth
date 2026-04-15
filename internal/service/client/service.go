package client

type service struct {
	cfg        Config
	txManager  txManager
	repository repository
	cache      cache
	outbox     outbox
}

func New(
	cfg Config,
	txManager txManager,
	repository repository,
	cache cache,
	outbox outbox,
) *service {
	return &service{
		cfg:        cfg,
		txManager:  txManager,
		repository: repository,
		cache:      cache,
		outbox:     outbox,
	}
}
