package client

type service struct {
	txManager  txManager
	repository repository
}

func New(txManager txManager, repository repository) *service {
	return &service{
		txManager:  txManager,
		repository: repository,
	}
}
