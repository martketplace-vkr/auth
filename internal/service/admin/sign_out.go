package admin

import (
	"context"
	"crypto/sha256"
	"encoding/base64"

	"github.com/martketplace-vkr/auth/internal/service/admin/dto"
)

func (s *service) SignOut(ctx context.Context, req dto.SignOutRequest) (err error) {
	hashBytes := sha256.Sum256([]byte(req.RefreshToken))
	refreshTokenHash := base64.URLEncoding.EncodeToString(hashBytes[:])

	err = s.cache.DeleteRefreshToken(ctx, refreshTokenHash)
	if err != nil {
		return err
	}

	return nil
}
