package repository

import "github.com/jmoiron/sqlx"

type IRepository interface {
}

type Repository struct {
	pqClient *sqlx.DB
}

func NewRepository(pqClient *sqlx.DB) IRepository {
	return &Repository{
		pqClient: pqClient,
	}
}
