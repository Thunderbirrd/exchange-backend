package service

import (
	"github.com/Thunderbirrd/exchange-backend/internal/repository"
	"github.com/Thunderbirrd/exchange-backend/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Api interface {
	CreateRequest(request models.Request) (int, error)
	GetRequests(data models.GetRequestsData) ([]models.Request, error)
	GetAllCurrencies() ([]models.Currency, error)
	GetAllAirports() ([]models.Airport, error)
}

type Service struct {
	Authorization
	Api
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Api:           NewApiService(repos.Api),
	}
}
