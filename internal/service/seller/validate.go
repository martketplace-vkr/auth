package seller

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
)

func (s *service) ValidateToken(
	ctx context.Context,
	accessToken string,
) (int64, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JwtSecret), nil
	})
	if err != nil {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := int64(claims["user_id"].(float64))

	return userID, nil
}
