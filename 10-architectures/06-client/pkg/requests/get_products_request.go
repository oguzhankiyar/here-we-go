package requests

import "client-sample/pkg/constants"

type GetProductsRequest struct {
	Offset int
	Limit  int
	Sort   string
	Order  constants.Order
}
