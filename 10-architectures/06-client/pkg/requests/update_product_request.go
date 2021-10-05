package requests

import "client-sample/pkg/constants"

type UpdateProductRequest struct {
	Id     string                  `json:"-"`
	Name   string                  `json:"name"`
	Price  float64                 `json:"price"`
	Status constants.ProductStatus `json:"status"`
}
