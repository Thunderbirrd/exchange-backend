package postgres

import (
	"fmt"
	"github.com/Thunderbirrd/exchange-backend/internal/dbo"
	"github.com/jmoiron/sqlx"
)

type ApiPostgres struct {
	db *sqlx.DB
}

func NewApiPostgres(db *sqlx.DB) *ApiPostgres {
	return &ApiPostgres{db: db}
}

func (r *ApiPostgres) CreateRequest(request dbo.Request) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (author_id, from_currency, to_currency, value_from, value_to, date_time, airport, status) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", requestsTable)
	row := r.db.QueryRow(query, request.AuthorId, request.FromCurrency, request.ToCurrency, request.ValueFrom, request.ValueTo, request.DateTime, request.Airport, request.Status)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
