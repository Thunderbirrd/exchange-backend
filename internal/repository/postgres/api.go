package postgres

import (
	"fmt"
	"github.com/Thunderbirrd/exchange-backend/internal/dbo"
	"github.com/Thunderbirrd/exchange-backend/pkg/models"
	"github.com/jmoiron/sqlx"
	"time"
)

type ApiPostgres struct {
	db *sqlx.DB
}

func NewApiPostgres(db *sqlx.DB) *ApiPostgres {
	return &ApiPostgres{db: db}
}

func (r *ApiPostgres) CreateRequest(request dbo.Request) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (author_id, from_currency, to_currency, value_from, value_to, date_time, airport) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", requestsTable)
	row := r.db.QueryRow(query, request.AuthorId, request.FromCurrency, request.ToCurrency, request.ValueFrom, request.ValueTo, request.DateTime, request.Airport)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ApiPostgres) GetRequests(from, to, airport string, valMin, valMax float32, dateTime time.Time) ([]dbo.Request, error) {
	var requests []dbo.Request
	query := fmt.Sprintf(`SELECT * FROM %s WHERE from_currency = $1 AND to_currency = $2 AND
								airport = $3 AND value_to BETWEEN $4 AND $5 AND date_time >= $6 ORDER BY date_time ASC`,
		requestsTable)

	if err := r.db.Select(&requests, query, from, to, airport, valMin, valMax, dateTime); err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *ApiPostgres) GetAllCurrencies() ([]models.Currency, error) {
	var currencies []models.Currency
	query := fmt.Sprintf("SELECT * FROM %s", currenciesTable)

	if err := r.db.Select(&currencies, query); err != nil {
		return nil, err
	}

	return currencies, nil
}

func (r *ApiPostgres) GetCurrencyByCode(code string) (models.Currency, error) {
	var currency models.Currency
	query := fmt.Sprintf("SELECT * FROM %s WHERE code = $1", currenciesTable)

	if err := r.db.Get(&currency, query, code); err != nil {
		return models.Currency{}, err
	}

	return currency, nil
}

func (r *ApiPostgres) GetAllAirports() ([]models.Airport, error) {
	var airports []models.Airport
	query := fmt.Sprintf("SELECT * FROM %s", airportsTable)

	if err := r.db.Select(&airports, query); err != nil {
		return nil, err
	}

	return airports, nil
}

func (r *ApiPostgres) GetAirportByCode(code string) (models.Airport, error) {
	var airport models.Airport
	query := fmt.Sprintf("SELECT * FROM %s WHERE code = $1", airportsTable)

	if err := r.db.Get(&airport, query, code); err != nil {
		return models.Airport{}, err
	}

	return airport, nil
}
