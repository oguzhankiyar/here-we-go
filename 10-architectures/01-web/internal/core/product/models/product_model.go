package models

import (
	"web-sample/internal/core/product/constants"
)

type ProductModel struct {
	Id        string                  `json:"id"`
	Name      string                  `json:"name"`
	Price     float64                 `json:"price"`
	Status    constants.ProductStatus `json:"status"`
	CreatedAt int64                   `json:"createdat"`
	UpdatedAt int64                   `json:"updatedAt"`
	DeletedAt int64                   `json:"deletedAt"`
}
