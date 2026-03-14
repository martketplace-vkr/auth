package client

type Config struct {
	RerfreshTokenTtlSeconds int64 `validate:"required" default:"300"`
}
