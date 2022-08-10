package main

import (
	"back/internal/app"
	"context"
	"flag"
	"log"
	
	"github.com/chapsuk/grace"
	"go.uber.org/zap"
	
	"back/internal/config"
	"back/pkg/logger"
	"back/pkg/postgres"
)

const (
	Name = "astrology"
)

// @title           API
// @version         1.0
// @description     API server

// @contact.name   Mekan Nepesov
// @contact.email  mnepesov.dev@gmail.com

// @host      localhost:8080
// @BasePath  /api
func main() {
	ctx := grace.ShutdownContext(context.Background())
	
	var logLevel, environment string
	flag.StringVar(&logLevel, "ll", "info", "logging level")
	flag.StringVar(&environment, "e", "", "environment")
	flag.Parse()
	
	cfg, err := config.NewAppConfig(environment)
	if err != nil {
		log.Fatal("error while read config: ", err.Error())
	}
	
	lg, err := logger.NewLogger(Name, environment, logLevel)
	if err != nil {
		log.Fatal("error while init logger", err.Error())
	}
	
	pqClient, err := postgres.NewPostgresDB(cfg.Postgres)
	if err != nil {
		lg.Fatal("error while connect to postgresql", zap.Error(err))
	}
	
	lg.Info(
		"flags",
		zap.String("name", Name),
		zap.String("environment", environment),
		zap.String("log_level", logLevel),
	)
	
	application := app.Application{
		Logger:         lg,
		Config:         cfg,
		PostgresClient: pqClient,
	}
	application.Run(ctx)
	application.Shutdown()
}
