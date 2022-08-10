package server

import (
	"context"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	errch "github.com/proxeter/errors-channel"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	
	_ "back/docs"
	"back/internal/config"
	"back/internal/nasa"
	"back/internal/repository"
	"back/internal/usecase"
)

type Server struct {
	Config *config.AppConfig
	Logger *zap.Logger
	
	postgresClient *sqlx.DB
}

func NewServer(
	ctx context.Context,
	logger *zap.Logger,
	config *config.AppConfig,
	pqClient *sqlx.DB,

) <-chan error {
	return errch.Register(func() error {
		return (&Server{
			Config:         config,
			Logger:         logger,
			postgresClient: pqClient,
		}).Start(ctx)
	})
}

func (s *Server) Start(ctx context.Context) error {
	repo := repository.NewRepository(s.postgresClient)
	n := nasa.NewNasa(s.Config.NASA.ApiKey)
	uc := usecase.NewUseCase(s.Logger, repo, n)
	
	router := s.initHTTPServer(uc)
	
	svr := http.Server{
		Handler: router,
		Addr:    ":" + s.Config.HTTP.Port,
	}
	
	s.Logger.Info(
		"Server running",
		zap.String("port", s.Config.HTTP.Port),
	)
	
	select {
	case err := <-errch.Register(svr.ListenAndServe):
		s.Logger.Info("Shutdown server", zap.String("by", "error"), zap.Error(err))
		return svr.Shutdown(ctx)
	case <-ctx.Done():
		s.Logger.Info("Shutdown server", zap.String("by", "context.Done"))
		return svr.Shutdown(ctx)
	}
}

func (s *Server) initHTTPServer(uc usecase.IUseCase) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	
	return router
}
