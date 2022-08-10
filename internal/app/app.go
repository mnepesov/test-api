package app

import (
	"context"
	"github.com/valyala/fasthttp"
	
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	
	"back/internal/app/server"
	"back/internal/config"
)

type Application struct {
	Config *config.AppConfig
	Logger *zap.Logger
	
	PostgresClient *sqlx.DB
	HttpClient     *fasthttp.Client
}

func (app *Application) Run(ctx context.Context) {
	httpServerErrCh := server.NewServer(
		ctx,
		app.Logger,
		app.Config,
		app.PostgresClient,
		app.HttpClient,
	)
	
	<-httpServerErrCh
}

func (app *Application) Shutdown() {
	app.Logger.Info("Shutdown database")
	_ = app.PostgresClient.Close()
	
	app.Logger.Info("Shutdown logger")
	_ = app.Logger.Sync()
}
