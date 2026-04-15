package seller

type Config struct {
	JwtSecret string `validate:"required"`
}
