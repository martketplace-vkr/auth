package app

import (
	"context"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	txmanager "github.com/avito-tech/go-transaction-manager/trm/manager"

	"github.com/martketplace-vkr/auth/config"
	"github.com/martketplace-vkr/auth/internal/app/cmp/server"
	clientRepository "github.com/martketplace-vkr/auth/internal/repository/pg/client"
	clientService "github.com/martketplace-vkr/auth/internal/service/client"
	clientTransport "github.com/martketplace-vkr/auth/internal/transport/grpc/v1/client"

	"github.com/martketplace-vkr/pkg/build"
	"github.com/martketplace-vkr/pkg/build/components/pgxsqlxcomponent"
)

func Run(ctx context.Context, cfg *config.Config) error {
	pg := pgxsqlxcomponent.New(cfg.Postgres)

	txManager, err := txmanager.New(trmsqlx.NewDefaultFactory(pg.DB))
	if err != nil {
		return err
	}

	clientRepo := clientRepository.New(pg.DB, trmsqlx.DefaultCtxGetter)
	clientServ := clientService.New(txManager, clientRepo)
	clientHandler := clientTransport.New(clientServ)

	grpcServer := server.New(
		cfg.Grpc,
		clientHandler,
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
