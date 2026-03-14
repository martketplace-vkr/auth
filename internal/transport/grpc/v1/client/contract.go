package client

import (
	"context"

	"github.com/martketplace-vkr/auth/internal/service/client/dto"
)

type (
	service interface {
		SignUp(ctx context.Context, req dto.SignUpRequest) (resp dto.SignUpResponse, err error)
	}
)
