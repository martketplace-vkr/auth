package app

import (
	"context"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	txmanager "github.com/avito-tech/go-transaction-manager/trm/manager"

	"github.com/martketplace-vkr/auth/config"
	"github.com/martketplace-vkr/auth/internal/app/cmp/server"
	adminRepository "github.com/martketplace-vkr/auth/internal/repository/pg/admin"
	clientRepository "github.com/martketplace-vkr/auth/internal/repository/pg/client"
	vendorRepository "github.com/martketplace-vkr/auth/internal/repository/pg/vendor"
	clientRedis "github.com/martketplace-vkr/auth/internal/repository/redis/client"
	adminService "github.com/martketplace-vkr/auth/internal/service/admin"
	clientService "github.com/martketplace-vkr/auth/internal/service/client"
	vendorService "github.com/martketplace-vkr/auth/internal/service/vendor"
	adminTransport "github.com/martketplace-vkr/auth/internal/transport/grpc/v1/admin"
	clientTransport "github.com/martketplace-vkr/auth/internal/transport/grpc/v1/client"
	vendorTransport "github.com/martketplace-vkr/auth/internal/transport/grpc/v1/vendor"

	"github.com/martketplace-vkr/pkg/build"
	"github.com/martketplace-vkr/pkg/build/components/pgxsqlxcomponent"
	"github.com/martketplace-vkr/pkg/build/components/rediscomponent"
)

func Run(ctx context.Context, cfg *config.Config) error {
	pg := pgxsqlxcomponent.New(cfg.Postgres)
	rd := rediscomponent.New(cfg.Redis)

	txManager, err := txmanager.New(trmsqlx.NewDefaultFactory(pg.DB))
	if err != nil {
		return err
	}

	clientRepo := clientRepository.New(pg.DB, trmsqlx.DefaultCtxGetter)
	adminRepo := adminRepository.New(pg.DB, trmsqlx.DefaultCtxGetter)
	vendorRepo := vendorRepository.New(pg.DB, trmsqlx.DefaultCtxGetter)
	clientCache := clientRedis.New(rd.Client)

	authCfg := clientService.Config{
		JwtSecret: cfg.AuthClientService.JwtSecret,
	}

	clientServ := clientService.New(cfg.AuthClientService, txManager, clientRepo, clientCache)
	adminServ := adminService.New(
		adminService.Config(authCfg),
		txManager,
		adminRepo,
		clientCache,
	)
	vendorServ := vendorService.New(
		vendorService.Config(authCfg),
		txManager,
		vendorRepo,
		clientCache,
	)

	clientHandler := clientTransport.New(clientServ)
	adminHandler := adminTransport.New(adminServ)
	vendorHandler := vendorTransport.New(vendorServ)

	grpcServer := server.New(
		cfg.Grpc,
		adminHandler,
		clientHandler,
		vendorHandler,
	)

	cmps := build.Components{
		pg,
		grpcServer,
	}

	app, err := build.NewApp(cmps)
	if err != nil {
		return err
	}

	return build.Run(ctx, app)
}
