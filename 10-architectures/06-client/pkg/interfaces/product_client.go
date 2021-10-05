package interfaces

import (
	"context"

	"client-sample/pkg/requests"
	"client-sample/pkg/responses"
)

type ProductClient interface {
	GetProducts(ctx context.Context, request *requests.GetProductsRequest) (*responses.GetProductsResponse, error)
	GetProductById(ctx context.Context, request *requests.GetProductByIdRequest) (*responses.GetProductByIdResponse, error)
	CreateProduct(ctx context.Context, request *requests.CreateProductRequest) (*responses.CreateProductResponse, error)
	UpdateProduct(ctx context.Context, request *requests.UpdateProductRequest) (*responses.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, request *requests.DeleteProductRequest) (*responses.DeleteProductResponse, error)
}
