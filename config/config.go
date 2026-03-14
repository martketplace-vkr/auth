package config

import (
	"github.com/martketplace-vkr/auth/internal/service/client"
	"github.com/martketplace-vkr/pkg/build/components/pgxsqlxcomponent"
	"github.com/martketplace-vkr/pkg/build/components/rediscomponent"
	"github.com/martketplace-vkr/pkg/server/grpc"
)

type Config struct {
	Grpc              grpc.Config             `validate:"required"`
	Postgres          pgxsqlxcomponent.Config `validate:"required"`
	Redis             rediscomponent.Config   `validate:"required"`
	AuthClientService client.Config           `validate:"required"`
}
