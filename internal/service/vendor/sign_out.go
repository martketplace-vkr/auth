package vendor

import (
	"context"
	"crypto/sha256"
	"encoding/base64"

	vendordto "github.com/martketplace-vkr/auth/internal/service/vendor_dto"
)

func (s *service) SignOut(ctx context.Context, req vendordto.SignOutRequest) (err error) {
	hashBytes := sha256.Sum256([]byte(req.RefreshToken))
	refreshTokenHash := base64.URLEncoding.EncodeToString(hashBytes[:])

	err = s.cache.DeleteRefreshToken(ctx, refreshTokenHash)
	if err != nil {
		return err
	}

	return nil
}
