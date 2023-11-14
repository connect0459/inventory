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

// BooksInfoAPIService is a service that implements the logic for the BooksInfoAPIServicer
// This service should implement the business logic for every endpoint for the BooksInfoAPI API.
// Include any external packages or services that will be required by this service.
type BooksInfoAPIService struct {
}

// NewBooksInfoAPIService creates a default api service
func NewBooksInfoAPIService() BooksInfoAPIServicer {
	return &BooksInfoAPIService{}
}

// Aefd1f8354eca94d504850ef936b3fb4 - Get a specific books_info by ID
func (s *BooksInfoAPIService) Aefd1f8354eca94d504850ef936b3fb4(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update Aefd1f8354eca94d504850ef936b3fb4 with the required logic for this service method.
	// Add api_books_info_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, BookInfo{}) or use other options such as http.Ok ...
	// return Response(200, BookInfo{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("Aefd1f8354eca94d504850ef936b3fb4 method not implemented")
}

// Call0c3283a323f1a4a3e81dbd5915be7518 - Create a new books_info
func (s *BooksInfoAPIService) Call0c3283a323f1a4a3e81dbd5915be7518(ctx context.Context, bookInfo BookInfo) (ImplResponse, error) {
	// TODO - update Call0c3283a323f1a4a3e81dbd5915be7518 with the required logic for this service method.
	// Add api_books_info_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Model0c3283a323f1a4a3e81dbd5915be7518201Response{}) or use other options such as http.Ok ...
	// return Response(201, Model0c3283a323f1a4a3e81dbd5915be7518201Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Call0c3283a323f1a4a3e81dbd5915be7518 method not implemented")
}

// Call2f9e356b188781492d6410beee2297c6 - Get a list of books_info
func (s *BooksInfoAPIService) Call2f9e356b188781492d6410beee2297c6(ctx context.Context) (ImplResponse, error) {
	// TODO - update Call2f9e356b188781492d6410beee2297c6 with the required logic for this service method.
	// Add api_books_info_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []BookInfo{}) or use other options such as http.Ok ...
	// return Response(200, []BookInfo{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Call2f9e356b188781492d6410beee2297c6 method not implemented")
}

// Call57b90f97a8a076e500862e557256c7aa - Update a specific books_info by ID
func (s *BooksInfoAPIService) Call57b90f97a8a076e500862e557256c7aa(ctx context.Context, id string, bookInfo BookInfo) (ImplResponse, error) {
	// TODO - update Call57b90f97a8a076e500862e557256c7aa with the required logic for this service method.
	// Add api_books_info_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Model0c3283a323f1a4a3e81dbd5915be7518201Response{}) or use other options such as http.Ok ...
	// return Response(200, Model0c3283a323f1a4a3e81dbd5915be7518201Response{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("Call57b90f97a8a076e500862e557256c7aa method not implemented")
}

// E98599424d7afa9076f6f9f4df52e4e4 - Delete a specific books_info by ID
func (s *BooksInfoAPIService) E98599424d7afa9076f6f9f4df52e4e4(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update E98599424d7afa9076f6f9f4df52e4e4 with the required logic for this service method.
	// Add api_books_info_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("E98599424d7afa9076f6f9f4df52e4e4 method not implemented")
}
