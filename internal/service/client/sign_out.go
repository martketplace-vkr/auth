package client

import (
	"context"

	"github.com/martketplace-vkr/auth/internal/service/client/dto"
)

func (s *service) SignOut(ctx context.Context, req dto.SignOutRequest) (err error) {
	err = s.cache.DeleteRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return err
	}

	return nil
}
