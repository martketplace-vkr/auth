package client

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) ValidateToken(
	ctx context.Context,
	accessToken string,
) (int64, string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JwtSecret), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	role, _ := claims["role"].(string)
	if role != "" && role != "client" {
		return 0, "", status.Error(codes.Unauthenticated, "invalid token role")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok || userIDFloat <= 0 {
		return 0, "", status.Error(codes.Unauthenticated, "invalid user id")
	}
	userID := int64(userIDFloat)
	login, _ := claims["login"].(string)
	user, err := s.repository.SelectUserByID(ctx, userID)
	if err != nil {
		return 0, "", status.Error(codes.Unauthenticated, "client not found")
	}
	if user.Status != "active" {
		return 0, "", status.Error(codes.PermissionDenied, "client account is blocked")
	}
	tokenVersion, _ := claims["token_version"].(float64)
	if int64(tokenVersion) != user.TokenVersion {
		return 0, "", status.Error(codes.Unauthenticated, "token was revoked")
	}
	if shouldRecord, cacheErr := s.cache.MarkActivity(ctx, userID); cacheErr == nil && shouldRecord {
		_ = s.repository.RecordActivity(ctx, userID)
	}

	return userID, login, nil
}
