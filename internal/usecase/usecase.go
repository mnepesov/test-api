package usecase

import (
	"back/internal/nasa"
	"go.uber.org/zap"
	
	"back/internal/repository"
)

type IUseCase interface {
}

type UseCase struct {
	logger *zap.Logger
	dbRepo repository.IRepository
	nasa   nasa.INasa
}

func NewUseCase(logger *zap.Logger, dbRepo repository.IRepository, n nasa.INasa) IUseCase {
	return &UseCase{
		logger: logger,
		dbRepo: dbRepo,
		nasa:   n,
	}
}
