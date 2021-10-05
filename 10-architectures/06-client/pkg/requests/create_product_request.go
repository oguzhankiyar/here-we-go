package requests

import "client-sample/pkg/constants"

type CreateProductRequest struct {
	Name   string                  `json:"name"`
	Price  float64                 `json:"price"`
	Status constants.ProductStatus `json:"status"`
}
