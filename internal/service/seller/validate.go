package seller

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
	if role != "" && role != "vendor" {
		return 0, "", status.Error(codes.Unauthenticated, "invalid token role")
	}

	userID := int64(claims["user_id"].(float64))
	login, _ := claims["login"].(string)

	return userID, login, nil
}
