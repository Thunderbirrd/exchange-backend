package repository

import (
	"github.com/Thunderbirrd/exchange-backend/internal/dbo"
	"github.com/Thunderbirrd/exchange-backend/internal/repository/postgres"
	"github.com/Thunderbirrd/exchange-backend/pkg/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Api interface {
	CreateRequest(request dbo.Request) (int, error)
	GetRequests(from, to, airport string, valMin, valMax float32, dateTime time.Time) ([]dbo.Request, error)
	GetAllCurrencies() ([]models.Currency, error)
	GetCurrencyByCode(code string) (models.Currency, error)
	GetAllAirports() ([]models.Airport, error)
	GetAirportByCode(code string) (models.Airport, error)
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
