/*
 * omu-rest
 *
 * Laravel製RESTful APIサーバー。著者: [connect0459](https://github.com/connect0459)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	BooksAPIService := openapi.NewBooksAPIService()
	BooksAPIController := openapi.NewBooksAPIController(BooksAPIService)

	BooksInfoAPIService := openapi.NewBooksInfoAPIService()
	BooksInfoAPIController := openapi.NewBooksInfoAPIController(BooksInfoAPIService)

	BooksStocksAPIService := openapi.NewBooksStocksAPIService()
	BooksStocksAPIController := openapi.NewBooksStocksAPIController(BooksStocksAPIService)

	ContactsAPIService := openapi.NewContactsAPIService()
	ContactsAPIController := openapi.NewContactsAPIController(ContactsAPIService)

	NewsAPIService := openapi.NewNewsAPIService()
	NewsAPIController := openapi.NewNewsAPIController(NewsAPIService)

	OrdersCustomersAPIService := openapi.NewOrdersCustomersAPIService()
	OrdersCustomersAPIController := openapi.NewOrdersCustomersAPIController(OrdersCustomersAPIService)

	OrdersItemsAPIService := openapi.NewOrdersItemsAPIService()
	OrdersItemsAPIController := openapi.NewOrdersItemsAPIController(OrdersItemsAPIService)

	OrdersPaymentsAPIService := openapi.NewOrdersPaymentsAPIService()
	OrdersPaymentsAPIController := openapi.NewOrdersPaymentsAPIController(OrdersPaymentsAPIService)

	TypesBranchesAPIService := openapi.NewTypesBranchesAPIService()
	TypesBranchesAPIController := openapi.NewTypesBranchesAPIController(TypesBranchesAPIService)

	TypesGenresAPIService := openapi.NewTypesGenresAPIService()
	TypesGenresAPIController := openapi.NewTypesGenresAPIController(TypesGenresAPIService)

	TypesOrdersStatesAPIService := openapi.NewTypesOrdersStatesAPIService()
	TypesOrdersStatesAPIController := openapi.NewTypesOrdersStatesAPIController(TypesOrdersStatesAPIService)

	TypesPaymentsAPIService := openapi.NewTypesPaymentsAPIService()
	TypesPaymentsAPIController := openapi.NewTypesPaymentsAPIController(TypesPaymentsAPIService)

	TypesReceivesAPIService := openapi.NewTypesReceivesAPIService()
	TypesReceivesAPIController := openapi.NewTypesReceivesAPIController(TypesReceivesAPIService)

	router := openapi.NewRouter(BooksAPIController, BooksInfoAPIController, BooksStocksAPIController, ContactsAPIController, NewsAPIController, OrdersCustomersAPIController, OrdersItemsAPIController, OrdersPaymentsAPIController, TypesBranchesAPIController, TypesGenresAPIController, TypesOrdersStatesAPIController, TypesPaymentsAPIController, TypesReceivesAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
