package vendor

type Config struct {
	JwtSecret string `validate:"required"`
}
