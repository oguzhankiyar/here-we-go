package entities

import (
	"time"

	"web-sample/internal/core/product/constants"
)

type ProductEntity struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Price     float64
	Status    constants.ProductStatus
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (ProductEntity) TableName() string {
	return "products"
}
