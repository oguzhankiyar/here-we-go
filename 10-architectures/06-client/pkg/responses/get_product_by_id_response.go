package responses

import "client-sample/pkg/models"

type GetProductByIdResponse struct {
	Product *models.ProductModel
}
