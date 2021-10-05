package queries

import "web-sample/internal/common/constants"

type FindProductsQuery struct {
	Offset int
	Limit  int
	Sort   string
	Order  constants.Order
}
