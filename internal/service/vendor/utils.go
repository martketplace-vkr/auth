package vendor

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) generateAccessToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(s.cfg.JwtSecret))
}

func generateRefreshToken() (token string, hash string, err error) {
	bytes := make([]byte, 32)

	_, err = rand.Read(bytes)
	if err != nil {
		return "", "", err
	}

	token = base64.URLEncoding.EncodeToString(bytes)

	hashBytes := sha256.Sum256([]byte(token))
	hash = base64.URLEncoding.EncodeToString(hashBytes[:])

	return token, hash, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}

func comparePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
