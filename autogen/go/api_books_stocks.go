/*
 * omu-rest
 *
 * Laravel製RESTful APIサーバー。著者: [connect0459](https://github.com/connect0459)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// BooksStocksAPIController binds http requests to an api service and writes the service results to the http response
type BooksStocksAPIController struct {
	service BooksStocksAPIServicer
	errorHandler ErrorHandler
}

// BooksStocksAPIOption for how the controller is set up.
type BooksStocksAPIOption func(*BooksStocksAPIController)

// WithBooksStocksAPIErrorHandler inject ErrorHandler into controller
func WithBooksStocksAPIErrorHandler(h ErrorHandler) BooksStocksAPIOption {
	return func(c *BooksStocksAPIController) {
		c.errorHandler = h
	}
}

// NewBooksStocksAPIController creates a default api controller
func NewBooksStocksAPIController(s BooksStocksAPIServicer, opts ...BooksStocksAPIOption) Router {
	controller := &BooksStocksAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the BooksStocksAPIController
func (c *BooksStocksAPIController) Routes() Routes {
	return Routes{
		"Call45a5cfdaf4972cd73313133082ca0d26": Route{
			strings.ToUpper("Delete"),
			"/api/books_stocks/{id}",
			c.Call45a5cfdaf4972cd73313133082ca0d26,
		},
		"Call46c90e16b416e254ed79948864ae87aa": Route{
			strings.ToUpper("Get"),
			"/api/books_stocks",
			c.Call46c90e16b416e254ed79948864ae87aa,
		},
		"D1c6d70f9a35791050b5da811f1abda1": Route{
			strings.ToUpper("Get"),
			"/api/books_stocks/{id}",
			c.D1c6d70f9a35791050b5da811f1abda1,
		},
		"Dc7eef146e782e14cddc076176ff04c4": Route{
			strings.ToUpper("Put"),
			"/api/books_stocks/{id}",
			c.Dc7eef146e782e14cddc076176ff04c4,
		},
		"Ecc649eea287e26856100c11f96bb3b9": Route{
			strings.ToUpper("Post"),
			"/api/books_stocks",
			c.Ecc649eea287e26856100c11f96bb3b9,
		},
	}
}

// Call45a5cfdaf4972cd73313133082ca0d26 - Delete a specific books_stocks by ID
func (c *BooksStocksAPIController) Call45a5cfdaf4972cd73313133082ca0d26(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.Call45a5cfdaf4972cd73313133082ca0d26(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Call46c90e16b416e254ed79948864ae87aa - Get a list of books_stocks
func (c *BooksStocksAPIController) Call46c90e16b416e254ed79948864ae87aa(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Call46c90e16b416e254ed79948864ae87aa(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// D1c6d70f9a35791050b5da811f1abda1 - Get a specific books_stocks by ID
func (c *BooksStocksAPIController) D1c6d70f9a35791050b5da811f1abda1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.D1c6d70f9a35791050b5da811f1abda1(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Dc7eef146e782e14cddc076176ff04c4 - Update a specific books_stocks by ID
func (c *BooksStocksAPIController) Dc7eef146e782e14cddc076176ff04c4(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	bookStockParam := BookStock{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bookStockParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertBookStockRequired(bookStockParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertBookStockConstraints(bookStockParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Dc7eef146e782e14cddc076176ff04c4(r.Context(), idParam, bookStockParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Ecc649eea287e26856100c11f96bb3b9 - Create a new books_stocks
func (c *BooksStocksAPIController) Ecc649eea287e26856100c11f96bb3b9(w http.ResponseWriter, r *http.Request) {
	bookStockParam := BookStock{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bookStockParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertBookStockRequired(bookStockParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertBookStockConstraints(bookStockParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Ecc649eea287e26856100c11f96bb3b9(r.Context(), bookStockParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}