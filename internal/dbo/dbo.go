package dbo

type Currency struct {
	Code      string  `db:"code"`
	RateToUsd float32 `db:"rate_to_usd"`
}

type Airport struct {
	Code            string `db:"code"`
	Country         string `db:"country"`
	City            string `db:"city"`
	MachineLocation string `db:"machine_location"`
}

type Request struct {
	Id           int     `db:"id"`
	AuthorId     int     `db:"author_id"`
	FromCurrency string  `db:"from_currency"`
	ToCurrency   string  `db:"to_currency"`
	ValueFrom    float32 `db:"value_from"`
	ValueTo      float32 `db:"value_to"`
	DateTime     string  `db:"date_time"`
	Airport      string  `db:"airport"`
	Status       string  `db:"status"`
}

type Exchange struct {
	Id              int    `db:"id"`
	Request         int    `db:"request_id"`
	AcceptorId      int    `db:"acceptor_id"`
	AuthorCode      string `db:"author_code"`
	AcceptorCode    string `db:"acceptor_code"`
	AuthorApprove   bool   `db:"author_approve"`
	AcceptorApprove bool   `db:"acceptor_approve"`
	ExpiredTime     string `db:"expired_time"`
}

type IdRequest struct {
	Id int `db:"id"`
}
