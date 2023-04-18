package dbo

import "time"

type Request struct {
	Id           int       `db:"id"`
	AuthorId     int       `db:"author_id"`
	FromCurrency string    `db:"from_currency"`
	ToCurrency   string    `db:"to_currency"`
	ValueFrom    float32   `db:"value_from"`
	ValueTo      float32   `db:"value_to"`
	DateTime     time.Time `db:"date_time"`
	Airport      string    `db:"airport"`
}

type Exchange struct {
	Id              int       `db:"id"`
	Request         int       `db:"request_id"`
	AuthorId        int       `db:"author_id"`
	AcceptorId      int       `db:"acceptor_id"`
	AuthorCode      string    `db:"author_code"`
	AcceptorCode    string    `db:"acceptor_code"`
	AuthorApprove   bool      `db:"author_approve"`
	AcceptorApprove bool      `db:"acceptor_approve"`
	ExpiredTime     time.Time `db:"expired_time"`
	Status          string    `db:"status"`
}

type IdRequest struct {
	Id int `db:"id"`
}

type UpdateExchangeInput struct {
	AuthorCode      *string
	AcceptorCode    *string
	AuthorApprove   *bool
	AcceptorApprove *bool
	Status          *string
}
