package service

import (
	"github.com/Thunderbirrd/exchange-backend/internal/dbo"
	"github.com/Thunderbirrd/exchange-backend/internal/repository"
	"github.com/Thunderbirrd/exchange-backend/pkg/models"
	"math/rand"
	"strconv"
	"time"
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

func (s *ApiService) CreateExchange(exchange models.Exchange) (int, error) {
	exchangeDbo := dbo.ExchangeToDbo(exchange)
	exchangeDbo.Status = string(models.Created)
	t, err := time.Parse(time.RFC3339, exchange.ExpiredTime)
	if err != nil {
		return 0, err
	}
	exchangeDbo.ExpiredTime = t

	return s.repo.CreateExchange(exchangeDbo)
}

func (s *ApiService) AcceptExchange(id int) error {
	status := string(models.InProgress)
	rand.Seed(time.Now().UnixNano())
	authorCode := strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
	acceptorCode := strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
	input := dbo.UpdateExchangeInput{Status: &status, AuthorCode: &authorCode, AcceptorCode: &acceptorCode}

	return s.repo.UpdateExchange(id, input)
}

func (s *ApiService) DeclineExchange(id int) error {
	status := string(models.Declined)
	input := dbo.UpdateExchangeInput{Status: &status}

	return s.repo.UpdateExchange(id, input)
}

func (s *ApiService) GetExchangeById(id int) (models.Exchange, error) {
	dboEx, err := s.repo.GetExchangeById(id)
	if err != nil {
		return models.Exchange{}, err
	}
	exchange := dbo.ExchangeToApi(dboEx)
	req, err := s.repo.GetRequestById(dboEx.Request)
	if err != nil {
		return models.Exchange{}, err
	}
	exchange.Request = dbo.RequestToApi(req)
	return exchange, nil
}

func (s *ApiService) GetUsersExchanges(userId int) ([]models.Exchange, error) {
	dboEx, err := s.repo.GetUsersExchanges(userId)
	if err != nil {
		return nil, err
	}

	var exchanges []models.Exchange
	var toApiEl models.Exchange
	for _, e := range dboEx {
		toApiEl = dbo.ExchangeToApi(e)
		req, err := s.repo.GetRequestById(e.Request)
		if err != nil {
			return nil, err
		}
		toApiEl.Request = dbo.RequestToApi(req)
		exchanges = append(exchanges, toApiEl)
	}

	return exchanges, nil
}

func (s *ApiService) GetAllCurrencies() ([]models.Currency, error) {
	return s.repo.GetAllCurrencies()
}

func (s *ApiService) GetAllAirports() ([]models.Airport, error) {
	return s.repo.GetAllAirports()
}
