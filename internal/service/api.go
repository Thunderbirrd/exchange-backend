package service

import (
	"github.com/Thunderbirrd/exchange-backend/internal/dbo"
	"github.com/Thunderbirrd/exchange-backend/internal/repository"
	"github.com/Thunderbirrd/exchange-backend/pkg/models"
)

type ApiService struct {
	repo repository.Api
}

func NewApiService(repo repository.Api) *ApiService {
	return &ApiService{repo: repo}
}

func (s *ApiService) CreateRequest(request models.Request) (int, error) {
	return s.repo.CreateRequest(dbo.RequestToDbo(request))
}
func (s *ApiService) GetRequests(data models.GetRequestsData) ([]models.Request, error) {
	valMin := data.Value * 0.9
	valMax := data.Value * 1.1

	dboReqs, err := s.repo.GetRequests(data.From, data.To, data.Airport, valMin, valMax, data.DateTime)
	if err != nil {
		return nil, err
	}
	var requests []models.Request
	var toApiEl models.Request
	for _, r := range dboReqs {
		toApiEl = dbo.RequestToApi(r)
		toApiEl.FromCurrency, err = s.repo.GetCurrencyByCode(r.FromCurrency)
		if err != nil {
			return nil, err
		}
		toApiEl.ToCurrency, err = s.repo.GetCurrencyByCode(r.ToCurrency)
		if err != nil {
			return nil, err
		}
		toApiEl.Airport, err = s.repo.GetAirportByCode(r.Airport)
		if err != nil {
			return nil, err
		}
		requests = append(requests, toApiEl)
	}

	return requests, nil
}

func (s *ApiService) GetAllCurrencies() ([]models.Currency, error) {
	return s.repo.GetAllCurrencies()
}

func (s *ApiService) GetAllAirports() ([]models.Airport, error) {
	return s.repo.GetAllAirports()
}
