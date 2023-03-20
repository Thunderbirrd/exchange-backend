package models

import "time"

type Status string

const (
	Created    Status = "created"
	InProgress Status = "in_progress"
	Finished   Status = "finished"
)

type User struct {
	Id       int     `json:"-" db:"id"`
	Name     string  `json:"name" binding:"required"`
	Surname  string  `json:"surname" binding:"required"`
	Login    string  `json:"login" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Rating   float32 `json:"rating"`
}

type Currency struct {
	Code      string  `json:"code" db:"code"`
	RateToUsd float32 `json:"rate_to_usd" db:"rate_to_usd"`
}

type Airport struct {
	Code            string `json:"code" db:"code"`
	Country         string `json:"country" db:"country"`
	City            string `json:"city" db:"city"`
	MachineLocation string `json:"machine_location" db:"machine_location"`
}

type Request struct {
	Id           int      `json:"id"`
	AuthorId     int      `json:"author_id"`
	FromCurrency Currency `json:"from_currency"`
	ToCurrency   Currency `json:"to_currency"`
	ValueFrom    float32  `json:"value_from"`
	ValueTo      float32  `json:"value_to"`
	DateTime     string   `json:"date_time"`
	Airport      Airport  `json:"airport"`
}

type Exchange struct {
	Id              int      `json:"id"`
	Request         *Request `json:"request"`
	AcceptorId      int      `json:"acceptor_id"`
	AuthorCode      string   `json:"author_code"`
	AcceptorCode    string   `json:"acceptor_code"`
	AuthorApprove   bool     `json:"author_approve"`
	AcceptorApprove bool     `json:"acceptor_approve"`
	ExpiredTime     string   `json:"expired_time"`
	Status          Status   `json:"status"`
}

type IdRequest struct {
	Id int `json:"id"`
}

type GetRequestsData struct {
	From     string    `json:"from" binding:"required"`
	To       string    `json:"to" binding:"required"`
	Airport  string    `json:"airport" binding:"required"`
	Value    float32   `json:"value" binding:"required"`
	DateTime time.Time `json:"date_time" binding:"required"`
}
