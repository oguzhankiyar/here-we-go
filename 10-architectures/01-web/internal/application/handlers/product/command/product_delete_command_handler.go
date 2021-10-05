package command

import (
	"context"

	"web-sample/internal/application/mappers/product"
	"web-sample/internal/core/product/commands"
	"web-sample/internal/core/product/errors"
	"web-sample/internal/core/product/repositories"
)

type DeleteProductCommandHandler struct {
	productRepository repositories.ProductRepository
	productMapper     *product.ProductMapper
}

func NewDeleteProductCommandHandler(productRepository repositories.ProductRepository, productMapper *product.ProductMapper) *DeleteProductCommandHandler {
	return &DeleteProductCommandHandler{
		productRepository: productRepository,
		productMapper:     productMapper,
	}
}

func (h DeleteProductCommandHandler) Handle(ctx context.Context, cmd *commands.DeleteProductCommand) error {
	entity, err := h.productRepository.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	if entity == nil {
		return errors.ProductNotFoundError
	}

	err = h.productRepository.Delete(ctx, entity.Id)

	if err != nil {
		return err
	}

	return nil
}
