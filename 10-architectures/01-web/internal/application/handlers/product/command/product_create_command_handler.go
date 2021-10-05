package command

import (
	"context"

	"web-sample/internal/application/mappers/product"
	"web-sample/internal/core/product/commands"
	"web-sample/internal/core/product/entities"
	"web-sample/internal/core/product/errors"
	"web-sample/internal/core/product/repositories"
)

type CreateProductCommandHandler struct {
	productRepository repositories.ProductRepository
	productMapper     *product.ProductMapper
}

func NewCreateProductCommandHandler(productRepository repositories.ProductRepository, productMapper *product.ProductMapper) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{
		productRepository: productRepository,
		productMapper:     productMapper,
	}
}

func (h CreateProductCommandHandler) Handle(ctx context.Context, cmd *commands.CreateProductCommand) (string, error) {
	entity, err := h.productRepository.FindByName(ctx, cmd.Name)

	if err != nil {
		return "", nil
	}

	if entity != nil {
		return "", errors.ProductAlreadyExistError
	}

	product := entities.ProductEntity{
		Name:   cmd.Name,
		Price:  cmd.Price,
		Status: cmd.Status,
	}

	id, err := h.productRepository.Create(ctx, &product)
	if err != nil {
		return "", nil
	}

	return id, nil
}
