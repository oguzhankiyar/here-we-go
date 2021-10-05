package command

import (
	"context"

	"web-sample/internal/application/mappers/product"
	"web-sample/internal/core/product/commands"
	"web-sample/internal/core/product/errors"
	"web-sample/internal/core/product/repositories"
)

type UpdateProductCommandHandler struct {
	productRepository repositories.ProductRepository
	productMapper     *product.ProductMapper
}

func NewUpdateProductCommandHandler(productRepository repositories.ProductRepository, productMapper *product.ProductMapper) *UpdateProductCommandHandler {
	return &UpdateProductCommandHandler{
		productRepository: productRepository,
		productMapper:     productMapper,
	}
}

func (h UpdateProductCommandHandler) Handle(ctx context.Context, cmd *commands.UpdateProductCommand) error {
	entity, err := h.productRepository.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	if entity == nil {
		return errors.ProductNotFoundError
	}

	entity.Name = cmd.Name
	entity.Price = cmd.Price
	entity.Status = cmd.Status

	err = h.productRepository.Update(ctx, entity)
	if err != nil {
		return err
	}

	return nil
}
