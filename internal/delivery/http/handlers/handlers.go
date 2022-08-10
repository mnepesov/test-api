package handlers

import (
	"go.uber.org/zap"
	
	"back/internal/usecase"
)

type Handler struct {
	logger  *zap.Logger
	useCase usecase.IUseCase
}

func NewHandler(logger *zap.Logger, useCase usecase.IUseCase) *Handler {
	return &Handler{
		logger:  logger,
		useCase: useCase,
	}
}
