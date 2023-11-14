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

// OrdersCustomersAPIController binds http requests to an api service and writes the service results to the http response
type OrdersCustomersAPIController struct {
	service OrdersCustomersAPIServicer
	errorHandler ErrorHandler
}

// OrdersCustomersAPIOption for how the controller is set up.
type OrdersCustomersAPIOption func(*OrdersCustomersAPIController)

// WithOrdersCustomersAPIErrorHandler inject ErrorHandler into controller
func WithOrdersCustomersAPIErrorHandler(h ErrorHandler) OrdersCustomersAPIOption {
	return func(c *OrdersCustomersAPIController) {
		c.errorHandler = h
	}
}

// NewOrdersCustomersAPIController creates a default api controller
func NewOrdersCustomersAPIController(s OrdersCustomersAPIServicer, opts ...OrdersCustomersAPIOption) Router {
	controller := &OrdersCustomersAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the OrdersCustomersAPIController
func (c *OrdersCustomersAPIController) Routes() Routes {
	return Routes{
		"A0a94057f0311e2193b5897f80221470": Route{
			strings.ToUpper("Put"),
			"/api/orders_customers/{id}",
			c.A0a94057f0311e2193b5897f80221470,
		},
		"Call4dd012aac3be308644f1e2e152f92c19": Route{
			strings.ToUpper("Get"),
			"/api/orders_customers/{id}",
			c.Call4dd012aac3be308644f1e2e152f92c19,
		},
		"Call5217a452519ec6ef916c6a3e2f4070dc": Route{
			strings.ToUpper("Get"),
			"/api/orders_customers",
			c.Call5217a452519ec6ef916c6a3e2f4070dc,
		},
		"Call5608eff2ef8984e0b7447765f27595aa": Route{
			strings.ToUpper("Delete"),
			"/api/orders_customers/{id}",
			c.Call5608eff2ef8984e0b7447765f27595aa,
		},
		"Call75e0485900f170d2424fdeadf59b0752": Route{
			strings.ToUpper("Post"),
			"/api/orders_customers",
			c.Call75e0485900f170d2424fdeadf59b0752,
		},
	}
}

// A0a94057f0311e2193b5897f80221470 - Update a specific orders_customers by ID
func (c *OrdersCustomersAPIController) A0a94057f0311e2193b5897f80221470(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	orderCustomerParam := OrderCustomer{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&orderCustomerParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOrderCustomerRequired(orderCustomerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertOrderCustomerConstraints(orderCustomerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.A0a94057f0311e2193b5897f80221470(r.Context(), idParam, orderCustomerParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Call4dd012aac3be308644f1e2e152f92c19 - Get a specific orders_customers by ID
func (c *OrdersCustomersAPIController) Call4dd012aac3be308644f1e2e152f92c19(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.Call4dd012aac3be308644f1e2e152f92c19(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Call5217a452519ec6ef916c6a3e2f4070dc - Get a list of orders_customers
func (c *OrdersCustomersAPIController) Call5217a452519ec6ef916c6a3e2f4070dc(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Call5217a452519ec6ef916c6a3e2f4070dc(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Call5608eff2ef8984e0b7447765f27595aa - Delete a specific orders_customers by ID
func (c *OrdersCustomersAPIController) Call5608eff2ef8984e0b7447765f27595aa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.Call5608eff2ef8984e0b7447765f27595aa(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Call75e0485900f170d2424fdeadf59b0752 - Create a new orders_customers
func (c *OrdersCustomersAPIController) Call75e0485900f170d2424fdeadf59b0752(w http.ResponseWriter, r *http.Request) {
	orderCustomerParam := OrderCustomer{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&orderCustomerParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOrderCustomerRequired(orderCustomerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertOrderCustomerConstraints(orderCustomerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Call75e0485900f170d2424fdeadf59b0752(r.Context(), orderCustomerParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}