package commands

import "web-sample/internal/core/product/constants"

type CreateProductCommand struct {
	Name   string                  `json:"name" validate:"required"`
	Price  float64                 `json:"price" validate:"required"`
	Status constants.ProductStatus `json:"status" validate:"required"`
}
