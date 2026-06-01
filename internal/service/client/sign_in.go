package client

import (
	"context"

	"github.com/martketplace-vkr/auth/internal/service/client/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) SignIn(ctx context.Context, req dto.SignInRequest) (resp dto.SignInResponse, err error) {
	user, err := s.repository.SelectUserByEmail(ctx, req.Email)
	if err != nil {
		return resp, err
	}

	err = comparePassword(user.PasswordHash, req.Password)
	if err != nil {
		return resp, err
	}
	if user.Status != "active" {
		return resp, status.Error(codes.PermissionDenied, "client account is blocked")
	}

	resp.AccessToken, err = s.generateAccessToken(user.ID, req.Email, user.TokenVersion)
	if err != nil {
		return resp, err
	}

	refreshToken, refreshHash, err := generateRefreshToken()
	if err != nil {
		return resp, err
	}

	resp.RefreshToken = refreshToken

	err = s.cache.SaveRefreshToken(ctx, dto.SaveRefreshTokenArgs{
		RefreshHash: refreshHash,
		UserID:      user.ID,
	})
	if err != nil {
		return resp, err
	}

	return resp, nil
}
