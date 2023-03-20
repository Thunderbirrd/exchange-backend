package handler

import (
	"github.com/Thunderbirrd/exchange-backend/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary CreateRequest
// @Security ApiKeyAuth
// @Tags api
// @Description create new request for exchange
// @ID create-request
// @Accept  json
// @Produce  json
// @Param input body models.Request true "request info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} models.errorResponse
// @Failure 500 {object} models.errorResponse
// @Failure default {object} models.errorResponse
// @Router /api/requests [post]
func (h *Handler) createRequest(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var request models.Request

	if err := c.BindJSON(&request); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	request.AuthorId = userId

	id, err := h.services.Api.CreateRequest(request)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary GetRequests
// @Security ApiKeyAuth
// @Tags api
// @Description get all request by specified params
// @ID get-requests
// @Accept  json
// @Produce  json
// @Param input body models.GetRequestsData true "request info"
// @Success 200 {integer} []models.Request
// @Failure 400,404 {object} models.errorResponse
// @Failure 500 {object} models.errorResponse
// @Failure default {object} models.errorResponse
// @Router /api/requests/get-all [post]
func (h *Handler) getRequests(c *gin.Context) {
	var data models.GetRequestsData

	if err := c.BindJSON(&data); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	requests, err := h.services.Api.GetRequests(data)
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, requests)
}

// @Summary GetCurrencies
// @Security ApiKeyAuth
// @Tags api
// @Description get currencies
// @ID get-currencies
// @Accept  json
// @Produce  json
// @Success 200 {integer} []models.Currency
// @Failure 400,404 {object} models.errorResponse
// @Failure 500 {object} models.errorResponse
// @Failure default {object} models.errorResponse
// @Router /api/currencies [get]
func (h *Handler) getAllCurrencies(c *gin.Context) {
	currencies, err := h.services.Api.GetAllCurrencies()
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, currencies)
}

// @Summary GetAirports
// @Security ApiKeyAuth
// @Tags api
// @Description get airports
// @ID get-airports
// @Accept  json
// @Produce  json
// @Success 200 {integer} []models.Airport
// @Failure 400,404 {object} models.errorResponse
// @Failure 500 {object} models.errorResponse
// @Failure default {object} models.errorResponse
// @Router /api/airports [get]
func (h *Handler) getAllAirports(c *gin.Context) {
	airports, err := h.services.Api.GetAllAirports()
	if err != nil {
		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, airports)
}
