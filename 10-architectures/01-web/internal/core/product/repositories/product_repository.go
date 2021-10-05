package repositories

import (
	"context"

	"web-sample/internal/common/constants"
	"web-sample/internal/core/product/entities"
)

type ProductRepository interface {
	Find(ctx context.Context, offset, limit int, sort string, order constants.Order) (*[]entities.ProductEntity, error)
	FindById(ctx context.Context, id string) (*entities.ProductEntity, error)
	FindByName(ctx context.Context, name string) (*entities.ProductEntity, error)
	Create(ctx context.Context, product *entities.ProductEntity) (string, error)
	Update(ctx context.Context, product *entities.ProductEntity) error
	Delete(ctx context.Context, id string) error
}
