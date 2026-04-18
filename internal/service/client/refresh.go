package client

import (
	"context"
	"crypto/sha256"
	"encoding/base64"

	"github.com/martketplace-vkr/auth/internal/service/client/dto"
)

func (s *service) Refresh(
	ctx context.Context,
	refreshToken string,
) (resp dto.RefreshResponse, err error) {

	hashBytes := sha256.Sum256([]byte(refreshToken))
	hash := base64.URLEncoding.EncodeToString(hashBytes[:])

	userID, err := s.cache.GetUserByRefreshHash(ctx, hash)
	if err != nil {
		return resp, err
	}

	user, err := s.repository.SelectUserByID(ctx, userID)
	if err != nil {
		return resp, err
	}

	resp.AccessToken, err = s.generateAccessToken(user.ID, user.Email)
	if err != nil {
		return resp, err
	}

	newRefreshToken, newHash, err := generateRefreshToken()
	if err != nil {
		return resp, err
	}

	resp.RefreshToken = newRefreshToken

	err = s.cache.SaveRefreshToken(ctx, dto.SaveRefreshTokenArgs{
		RefreshHash: newHash,
		UserID:      userID,
	})
	if err != nil {
		return resp, err
	}

	return resp, nil
}
