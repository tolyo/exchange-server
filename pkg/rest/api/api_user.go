/*
 * OPEN OUTCRY API
 *
 * # Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// UserAPIController binds http requests to an api service and writes the service results to the http response
type UserAPIController struct {
	service      UserAPIServicer
	errorHandler ErrorHandler
}

// UserAPIOption for how the controller is set up.
type UserAPIOption func(*UserAPIController)

// WithUserAPIErrorHandler inject ErrorHandler into controller
func WithUserAPIErrorHandler(h ErrorHandler) UserAPIOption {
	return func(c *UserAPIController) {
		c.errorHandler = h
	}
}

// NewUserAPIController creates a default api controller
func NewUserAPIController(s UserAPIServicer, opts ...UserAPIOption) Router {
	controller := &UserAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UserAPIController
func (c *UserAPIController) Routes() Routes {
	return Routes{
		"CreateTrade": Route{
			strings.ToUpper("Post"),
			"/trade_orders/{trading_account_id}",
			c.CreateTrade,
		},
		"DeleteTradeById": Route{
			strings.ToUpper("Delete"),
			"/trades/{trading_account_id}/id/{trade_id}",
			c.DeleteTradeById,
		},
		"GetBookOrders": Route{
			strings.ToUpper("Get"),
			"/book_orders/{trading_account_id}",
			c.GetBookOrders,
		},
		"GetPaymentAccounts": Route{
			strings.ToUpper("Get"),
			"/payment-accounts/{app_entity_id}",
			c.GetPaymentAccounts,
		},
		"GetTradeById": Route{
			strings.ToUpper("Get"),
			"/trades/{trading_account_id}/id/{trade_id}",
			c.GetTradeById,
		},
		"GetTradeOrders": Route{
			strings.ToUpper("Get"),
			"/trade_orders/{trading_account_id}",
			c.GetTradeOrders,
		},
		"GetTrades": Route{
			strings.ToUpper("Get"),
			"/trades/{trading_account_id}",
			c.GetTrades,
		},
		"GetTradingAccount": Route{
			strings.ToUpper("Get"),
			"/trading_account/{trading_account_id}",
			c.GetTradingAccount,
		},
	}
}

// CreateTrade - Create trade order
func (c *UserAPIController) CreateTrade(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tradingAccountIdParam := params["trading_account_id"]
	if tradingAccountIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"trading_account_id"}, nil)
		return
	}
	createTradeRequestParam := CreateTradeRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&createTradeRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateTradeRequestRequired(createTradeRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCreateTradeRequestConstraints(createTradeRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateTrade(r.Context(), tradingAccountIdParam, createTradeRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteTradeById - Cancel trade
func (c *UserAPIController) DeleteTradeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tradingAccountIdParam := params["trading_account_id"]
	if tradingAccountIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"trading_account_id"}, nil)
		return
	}
	tradeIdParam := params["trade_id"]
	if tradeIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"trade_id"}, nil)
		return
	}
	result, err := c.service.DeleteTradeById(r.Context(), tradingAccountIdParam, tradeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetBookOrders - Get book orders
func (c *UserAPIController) GetBookOrders(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tradingAccountIdParam := params["trading_account_id"]
	if tradingAccountIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"trading_account_id"}, nil)
		return
	}
	result, err := c.service.GetBookOrders(r.Context(), tradingAccountIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPaymentAccounts - Get payment accounts
func (c *UserAPIController) GetPaymentAccounts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appEntityIdParam := params["app_entity_id"]
	if appEntityIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_entity_id"}, nil)
		return
	}
	result, err := c.service.GetPaymentAccounts(r.Context(), appEntityIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetTradeById - Get trade
func (c *UserAPIController) GetTradeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tradingAccountIdParam := params["trading_account_id"]
	if tradingAccountIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"trading_account_id"}, nil)
		return
	}
	tradeIdParam := params["trade_id"]
	if tradeIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"trade_id"}, nil)
		return
	}
	result, err := c.service.GetTradeById(r.Context(), tradingAccountIdParam, tradeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetTradeOrders - Get trade orders
func (c *UserAPIController) GetTradeOrders(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tradingAccountIdParam := params["trading_account_id"]
	if tradingAccountIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"trading_account_id"}, nil)
		return
	}
	result, err := c.service.GetTradeOrders(r.Context(), tradingAccountIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetTrades - Trades list
func (c *UserAPIController) GetTrades(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tradingAccountIdParam := params["trading_account_id"]
	if tradingAccountIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"trading_account_id"}, nil)
		return
	}
	result, err := c.service.GetTrades(r.Context(), tradingAccountIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetTradingAccount - Get trading account
func (c *UserAPIController) GetTradingAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tradingAccountIdParam := params["trading_account_id"]
	if tradingAccountIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"trading_account_id"}, nil)
		return
	}
	result, err := c.service.GetTradingAccount(r.Context(), tradingAccountIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
