package query

import (
	"context"

	"web-sample/internal/application/mappers/product"
	"web-sample/internal/core/product/models"
	"web-sample/internal/core/product/queries"
	"web-sample/internal/core/product/repositories"
)

type FindProductsQueryHandler struct {
	productRepository repositories.ProductRepository
	productMapper     *product.ProductMapper
}

func NewFindProductsQueryHandler(productRepository repositories.ProductRepository, productMapper *product.ProductMapper) *FindProductsQueryHandler {
	return &FindProductsQueryHandler{
		productRepository: productRepository,
		productMapper:     productMapper,
	}
}

func (h FindProductsQueryHandler) Handle(ctx context.Context, qry *queries.FindProductsQuery) (*[]models.ProductModel, error) {
	entities, err := h.productRepository.Find(ctx, qry.Offset, qry.Limit, qry.Sort, qry.Order)

	if err != nil {
		return nil, err
	}

	return h.productMapper.MapMany(entities), nil
}
