package app

import (
	"context"
	"github.com/Kokkibegushidoktor/test1/internal/app/http"
	"github.com/Kokkibegushidoktor/test1/internal/app/http/handlers"
	"github.com/Kokkibegushidoktor/test1/internal/clients"
	"github.com/Kokkibegushidoktor/test1/internal/config"
	"github.com/Kokkibegushidoktor/test1/internal/repository"
	"github.com/Kokkibegushidoktor/test1/internal/service"
	"github.com/Kokkibegushidoktor/test1/internal/utils"
)

func Run(ctx context.Context, cfg *config.Config) error {
	mng := clients.NewMongoClient(ctx, cfg)
	db := mng.Database(cfg.MngDbName)

	repos := repository.NewRepositories(db)

	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	services.Tickers.InitUpdateWorker(ctx)

	handler := handlers.New(services)

	httpServer := http.New(cfg, handler)
	httpServer.Start()

	utils.GracefulShutdown()

	return nil
}
