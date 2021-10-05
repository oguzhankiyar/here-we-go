package query

import (
	"context"

	"web-sample/internal/application/mappers/product"
	"web-sample/internal/core/product/errors"
	"web-sample/internal/core/product/models"
	"web-sample/internal/core/product/queries"
	"web-sample/internal/core/product/repositories"
)

type FindProductByIdQueryHandler struct {
	productRepository repositories.ProductRepository
	productMapper     *product.ProductMapper
}

func NewFindProductByIdQueryHandler(productRepository repositories.ProductRepository, productMapper *product.ProductMapper) *FindProductByIdQueryHandler {
	return &FindProductByIdQueryHandler{
		productRepository: productRepository,
		productMapper:     productMapper,
	}
}

func (h FindProductByIdQueryHandler) Handle(ctx context.Context, qry *queries.FindProductByIdQuery) (*models.ProductModel, error) {
	entity, err := h.productRepository.FindById(ctx, qry.Id)
	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, errors.ProductNotFoundError
	}

	return h.productMapper.MapOne(entity), nil
}
