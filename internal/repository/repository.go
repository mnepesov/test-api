package repository

import (
	"back/internal/domains"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type IRepository interface {
	AddAPOD(ctx context.Context, apod domains.NasaAPOD) error
	GetAPODs(ctx context.Context) ([]domains.APOD, error)
	GetAPODsByDate(ctx context.Context, date string) ([]domains.APOD, error)
}

type Repository struct {
	pqClient *sqlx.DB
}

func NewRepository(pqClient *sqlx.DB) IRepository {
	return &Repository{
		pqClient: pqClient,
	}
}

func (r *Repository) AddAPOD(ctx context.Context, apod domains.NasaAPOD) error {
	data, err := json.Marshal(apod)
	if err != nil {
		return err
	}
	
	query := "insert into pictures(apod) values ($1)"
	_, err = r.pqClient.ExecContext(ctx, query, string(data))
	return err
}

func (r *Repository) GetAPODs(ctx context.Context) ([]domains.APOD, error) {
	query := "select id, apod, created_at from pictures"
	rows, err := r.pqClient.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	
	var apods []domains.APOD
	for rows.Next() {
		var apod domains.APOD
		var data string
		if err := rows.Scan(&apod.ID, &data, &apod.CreatedAt); err != nil {
			continue
		}
		
		err = json.Unmarshal([]byte(data), &apod.NasaAPOD)
		if err != nil {
			continue
		}
		
		apods = append(apods, apod)
	}
	
	return apods, nil
}

func (r *Repository) GetAPODsByDate(ctx context.Context, date string) ([]domains.APOD, error) {
	query := "select id, apod, created_at from pictures where created_at = $1"
	rows, err := r.pqClient.QueryContext(ctx, query, date)
	if err != nil {
		return nil, err
	}
	
	var apods []domains.APOD
	for rows.Next() {
		var apod domains.APOD
		var data string
		if err := rows.Scan(&apod.ID, &data, &apod.CreatedAt); err != nil {
			continue
		}
		
		err = json.Unmarshal([]byte(data), &apod.NasaAPOD)
		if err != nil {
			continue
		}
		
		apods = append(apods, apod)
	}
	
	return apods, nil
}
