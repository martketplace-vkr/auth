package admin

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
