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

// OrdersItemsAPIController binds http requests to an api service and writes the service results to the http response
type OrdersItemsAPIController struct {
	service OrdersItemsAPIServicer
	errorHandler ErrorHandler
}

// OrdersItemsAPIOption for how the controller is set up.
type OrdersItemsAPIOption func(*OrdersItemsAPIController)

// WithOrdersItemsAPIErrorHandler inject ErrorHandler into controller
func WithOrdersItemsAPIErrorHandler(h ErrorHandler) OrdersItemsAPIOption {
	return func(c *OrdersItemsAPIController) {
		c.errorHandler = h
	}
}

// NewOrdersItemsAPIController creates a default api controller
func NewOrdersItemsAPIController(s OrdersItemsAPIServicer, opts ...OrdersItemsAPIOption) Router {
	controller := &OrdersItemsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the OrdersItemsAPIController
func (c *OrdersItemsAPIController) Routes() Routes {
	return Routes{
		"B2909cc3849277b1d5345ccb4205c71d": Route{
			strings.ToUpper("Post"),
			"/api/orders_items",
			c.B2909cc3849277b1d5345ccb4205c71d,
		},
		"Call0b2a7a56e7d5c2c49b7d11bdb0f48f5c": Route{
			strings.ToUpper("Get"),
			"/api/orders_items/{id}",
			c.Call0b2a7a56e7d5c2c49b7d11bdb0f48f5c,
		},
		"Call6d397dbc3afd3a3f44e4cc08820e0921": Route{
			strings.ToUpper("Delete"),
			"/api/orders_items/{id}",
			c.Call6d397dbc3afd3a3f44e4cc08820e0921,
		},
		"Call89fde6ad08076b992c521fb7c680e29e": Route{
			strings.ToUpper("Get"),
			"/api/orders_items",
			c.Call89fde6ad08076b992c521fb7c680e29e,
		},
		"Call8bee2e95aba24db3af9b2bc7b5a53f0d": Route{
			strings.ToUpper("Put"),
			"/api/orders_items/{id}",
			c.Call8bee2e95aba24db3af9b2bc7b5a53f0d,
		},
	}
}

// B2909cc3849277b1d5345ccb4205c71d - Create a new orders_items
func (c *OrdersItemsAPIController) B2909cc3849277b1d5345ccb4205c71d(w http.ResponseWriter, r *http.Request) {
	orderItemParam := OrderItem{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&orderItemParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOrderItemRequired(orderItemParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertOrderItemConstraints(orderItemParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.B2909cc3849277b1d5345ccb4205c71d(r.Context(), orderItemParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Call0b2a7a56e7d5c2c49b7d11bdb0f48f5c - Get a specific orders_items by ID
func (c *OrdersItemsAPIController) Call0b2a7a56e7d5c2c49b7d11bdb0f48f5c(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.Call0b2a7a56e7d5c2c49b7d11bdb0f48f5c(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Call6d397dbc3afd3a3f44e4cc08820e0921 - Delete a specific orders_items by ID
func (c *OrdersItemsAPIController) Call6d397dbc3afd3a3f44e4cc08820e0921(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.Call6d397dbc3afd3a3f44e4cc08820e0921(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Call89fde6ad08076b992c521fb7c680e29e - Get a list of orders_items
func (c *OrdersItemsAPIController) Call89fde6ad08076b992c521fb7c680e29e(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.Call89fde6ad08076b992c521fb7c680e29e(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// Call8bee2e95aba24db3af9b2bc7b5a53f0d - Update a specific orders_items by ID
func (c *OrdersItemsAPIController) Call8bee2e95aba24db3af9b2bc7b5a53f0d(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	orderItemParam := OrderItem{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&orderItemParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOrderItemRequired(orderItemParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertOrderItemConstraints(orderItemParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Call8bee2e95aba24db3af9b2bc7b5a53f0d(r.Context(), idParam, orderItemParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
