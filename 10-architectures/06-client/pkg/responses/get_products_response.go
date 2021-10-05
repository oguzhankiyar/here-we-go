package responses

import "client-sample/pkg/models"

type GetProductsResponse struct {
	Products *[]models.ProductModel
}
