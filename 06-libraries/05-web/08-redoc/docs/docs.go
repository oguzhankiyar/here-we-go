// Package classification of Item API
//
// Documentation for Item API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package docs

import (
	"fmt"
	"redoc-sample/models"
)

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of items
// swagger:response itemsResponse
type itemsResponseWrapper struct {
	// All current items
	// in: body
	Body []models.Item
}

// Data structure representing a single item
// swagger:response itemResponse
type itemResponseWrapper struct {
	// Newly created item
	// in: body
	Body models.Item
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateItem createItem
type itemParamsWrapper struct {
	// Item data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body models.Item
}

// swagger:parameters updateItem
type itemIDParamsWrapper struct {
	// The id of the item for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// ErrInvalidItemPath is an error message when the item path is not valid
var ErrInvalidItemPath = fmt.Errorf("invalid path, path should be /items/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
