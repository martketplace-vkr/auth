package client

type Config struct {
	JwtSecret string `validate:"required"`
}
