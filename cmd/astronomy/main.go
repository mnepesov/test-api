package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"
	"time"
	
	"github.com/chapsuk/grace"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	
	"back/internal/app"
	"back/internal/config"
	"back/pkg/logger"
	"back/pkg/postgres"
)

const (
	Name = "astronomy"
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
	
	httpClient := &fasthttp.Client{
		Name:                "Go-http-client/1.1",
		MaxConnsPerHost:     32,
		MaxIdleConnDuration: 60 * time.Second,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
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
		HttpClient:     httpClient,
	}
	application.Run(ctx)
	application.Shutdown()
}
