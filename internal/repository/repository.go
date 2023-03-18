package repository

import (
	"github.com/Thunderbirrd/exchange-backend/internal/dbo"
	"github.com/Thunderbirrd/exchange-backend/internal/repository/postgres"
	"github.com/Thunderbirrd/exchange-backend/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Api interface {
	CreateRequest(request dbo.Request) (int, error)
}

type Repository struct {
	Authorization
	Api
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		Api:           postgres.NewApiPostgres(db),
	}
}
