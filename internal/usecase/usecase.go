package usecase

import (
	"context"
	"errors"
	
	"go.uber.org/zap"
	
	"back/internal/domains"
	"back/internal/nasa"
	"back/internal/repository"
)

var (
	ErrGetAPOD = errors.New("unable to get APOD")
	ErrAddApod = errors.New("unable to add APOD")
)

type IUseCase interface {
	GetAPODFromNasa(ctx context.Context) (domains.NasaAPOD, error)
	AddAPOD(ctx context.Context, apod domains.NasaAPOD) error
	GetAPODs(ctx context.Context) ([]domains.APOD, error)
	GetAPODsByDate(ctx context.Context, date string) ([]domains.APOD, error)
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

func (uc *UseCase) GetAPODFromNasa(ctx context.Context) (domains.NasaAPOD, error) {
	apod, err := uc.nasa.GetAPOD()
	if err != nil {
		uc.logger.Error("usecase.GetAPODFromNasa", zap.Error(err))
		return domains.NasaAPOD{}, ErrGetAPOD
	}
	
	a := domains.NasaAPOD{
		Title:          apod.Title,
		Copyright:      apod.Copyright,
		Explanation:    apod.Explanation,
		HdUrl:          apod.HdUrl,
		URL:            apod.URL,
		MediaType:      apod.MediaType,
		ServiceVersion: apod.ServiceVersion,
		Date:           apod.Date,
	}
	
	return a, nil
}

func (uc *UseCase) AddAPOD(ctx context.Context, apod domains.NasaAPOD) error {
	err := uc.dbRepo.AddAPOD(ctx, apod)
	if err != nil {
		uc.logger.Error("usecase.AddAPOD", zap.Error(err))
		return ErrAddApod
	}
	
	return nil
}

func (uc *UseCase) GetAPODs(ctx context.Context) ([]domains.APOD, error) {
	apods, err := uc.dbRepo.GetAPODs(ctx)
	if err != nil {
		uc.logger.Error("usecase.GetAPODs", zap.Error(err))
		return nil, ErrGetAPOD
	}
	
	return apods, nil
}

func (uc *UseCase) GetAPODsByDate(ctx context.Context, date string) ([]domains.APOD, error) {
	apods, err := uc.dbRepo.GetAPODsByDate(ctx, date)
	if err != nil {
		uc.logger.Error("usecase.GetAPODsByDate", zap.Error(err))
		return nil, ErrGetAPOD
	}
	
	return apods, nil
}
