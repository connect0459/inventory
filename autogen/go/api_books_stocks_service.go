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
	"context"
	"net/http"
	"errors"
)

// BooksStocksAPIService is a service that implements the logic for the BooksStocksAPIServicer
// This service should implement the business logic for every endpoint for the BooksStocksAPI API.
// Include any external packages or services that will be required by this service.
type BooksStocksAPIService struct {
}

// NewBooksStocksAPIService creates a default api service
func NewBooksStocksAPIService() BooksStocksAPIServicer {
	return &BooksStocksAPIService{}
}

// Call45a5cfdaf4972cd73313133082ca0d26 - Delete a specific books_stocks by ID
func (s *BooksStocksAPIService) Call45a5cfdaf4972cd73313133082ca0d26(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update Call45a5cfdaf4972cd73313133082ca0d26 with the required logic for this service method.
	// Add api_books_stocks_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("Call45a5cfdaf4972cd73313133082ca0d26 method not implemented")
}

// Call46c90e16b416e254ed79948864ae87aa - Get a list of books_stocks
func (s *BooksStocksAPIService) Call46c90e16b416e254ed79948864ae87aa(ctx context.Context) (ImplResponse, error) {
	// TODO - update Call46c90e16b416e254ed79948864ae87aa with the required logic for this service method.
	// Add api_books_stocks_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []BookStock{}) or use other options such as http.Ok ...
	// return Response(200, []BookStock{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Call46c90e16b416e254ed79948864ae87aa method not implemented")
}

// D1c6d70f9a35791050b5da811f1abda1 - Get a specific books_stocks by ID
func (s *BooksStocksAPIService) D1c6d70f9a35791050b5da811f1abda1(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update D1c6d70f9a35791050b5da811f1abda1 with the required logic for this service method.
	// Add api_books_stocks_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, BookStock{}) or use other options such as http.Ok ...
	// return Response(200, BookStock{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("D1c6d70f9a35791050b5da811f1abda1 method not implemented")
}

// Dc7eef146e782e14cddc076176ff04c4 - Update a specific books_stocks by ID
func (s *BooksStocksAPIService) Dc7eef146e782e14cddc076176ff04c4(ctx context.Context, id string, bookStock BookStock) (ImplResponse, error) {
	// TODO - update Dc7eef146e782e14cddc076176ff04c4 with the required logic for this service method.
	// Add api_books_stocks_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Ecc649eea287e26856100c11f96bb3b9201Response{}) or use other options such as http.Ok ...
	// return Response(200, Ecc649eea287e26856100c11f96bb3b9201Response{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("Dc7eef146e782e14cddc076176ff04c4 method not implemented")
}

// Ecc649eea287e26856100c11f96bb3b9 - Create a new books_stocks
func (s *BooksStocksAPIService) Ecc649eea287e26856100c11f96bb3b9(ctx context.Context, bookStock BookStock) (ImplResponse, error) {
	// TODO - update Ecc649eea287e26856100c11f96bb3b9 with the required logic for this service method.
	// Add api_books_stocks_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Ecc649eea287e26856100c11f96bb3b9201Response{}) or use other options such as http.Ok ...
	// return Response(201, Ecc649eea287e26856100c11f96bb3b9201Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Ecc649eea287e26856100c11f96bb3b9 method not implemented")
}