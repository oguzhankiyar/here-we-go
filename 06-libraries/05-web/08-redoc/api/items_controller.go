package api

import (
	"net/http"
)

// ItemsController for getting and updating items
type ItemsController struct {

}

// NewItemsController returns new item controller
func NewItemsController() *ItemsController {
	return &ItemsController{}
}

// swagger:route GET /items items getAll
// Return a list of items from the database
// responses:
//	200: itemsResponse

// GetAll handles GET requests and returns all current items
func (c *ItemsController) GetAll(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("[]"))
}

// swagger:route GET /items/{id} items getOne
// Return a list of items from the database
// responses:
//	200: itemResponse
//	404: errorResponse

// GetOne handles GET requests
func (c *ItemsController) GetOne(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("{}"))
}

// swagger:route POST /items items createItem
// Create a new item
//
// responses:
//	200: itemResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new items
func (c *ItemsController) Create(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNoContent)
}

// swagger:route PUT /items items updateItem
// Update a item details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update items
func (c *ItemsController) Update(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNoContent)
}

// swagger:route DELETE /items/{id} items deleteItem
// Delete a item
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes item from the database
func (c *ItemsController) Delete(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNoContent)
}
