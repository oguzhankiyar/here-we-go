package product

import (
	"web-sample/internal/application/handlers/product/command"
	"web-sample/internal/application/handlers/product/query"
	"web-sample/internal/application/mappers/product"
	"web-sample/internal/core/product/repositories"
)

type ProductHandler struct {
	FindProducts    *query.FindProductsQueryHandler
	FindProductById *query.FindProductByIdQueryHandler
	CreateProduct   *command.CreateProductCommandHandler
	UpdateProduct   *command.UpdateProductCommandHandler
	DeleteProduct   *command.DeleteProductCommandHandler
}

func NewProductHandler(productRepository repositories.ProductRepository, productMapper *product.ProductMapper) *ProductHandler {
	return &ProductHandler{
		FindProducts:    query.NewFindProductsQueryHandler(productRepository, productMapper),
		FindProductById: query.NewFindProductByIdQueryHandler(productRepository, productMapper),
		CreateProduct:   command.NewCreateProductCommandHandler(productRepository, productMapper),
		UpdateProduct:   command.NewUpdateProductCommandHandler(productRepository, productMapper),
		DeleteProduct:   command.NewDeleteProductCommandHandler(productRepository, productMapper),
	}
}
