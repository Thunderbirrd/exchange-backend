package dbo

import (
	"github.com/Thunderbirrd/exchange-backend/pkg/models"
	"log"
	"time"
)

func RequestToDbo(request models.Request) Request {
	var dbo = Request{
		Id:           request.Id,
		AuthorId:     request.AuthorId,
		FromCurrency: request.FromCurrency.Code,
		ToCurrency:   request.ToCurrency.Code,
		ValueFrom:    request.ValueFrom,
		ValueTo:      request.ValueTo,
		Airport:      request.Airport.Code,
	}
	t, err := time.Parse(time.RFC3339, request.DateTime)
	if err != nil {
		log.Fatalln(err.Error())
	}
	dbo.DateTime = t
	return dbo
}

func RequestToApi(request Request) models.Request {
	return models.Request{
		Id:        request.Id,
		AuthorId:  request.AuthorId,
		ValueTo:   request.ValueTo,
		ValueFrom: request.ValueFrom,
		DateTime:  request.DateTime.Format(time.RFC3339),
	}
}

func ExchangeToDbo(exchange models.Exchange) Exchange {
	return Exchange{
		Id:              exchange.Id,
		Request:         exchange.Request.Id,
		AuthorId:        exchange.AuthorId,
		AcceptorId:      exchange.AcceptorId,
		AcceptorCode:    exchange.AcceptorCode,
		AuthorCode:      exchange.AuthorCode,
		AuthorApprove:   exchange.AuthorApprove,
		AcceptorApprove: exchange.AcceptorApprove,
	}
}

func ExchangeToApi(exchange Exchange) models.Exchange {
	return models.Exchange{
		Id:              exchange.Id,
		AuthorId:        exchange.AuthorId,
		AcceptorId:      exchange.AcceptorId,
		AcceptorCode:    exchange.AcceptorCode,
		AuthorCode:      exchange.AuthorCode,
		AuthorApprove:   exchange.AuthorApprove,
		AcceptorApprove: exchange.AcceptorApprove,
		Status:          models.Status(exchange.Status),
		ExpiredTime:     exchange.ExpiredTime.Format(time.RFC3339),
	}
}
