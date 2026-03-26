package admin

type Config struct {
	JwtSecret string `validate:"required"`
}
