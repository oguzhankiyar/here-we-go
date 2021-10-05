package product

import (
	"web-sample/internal/core/product/entities"
	"web-sample/internal/core/product/models"
)

type ProductMapper struct {
}

func NewProductMapper() *ProductMapper {
	return &ProductMapper{}
}

func (m ProductMapper) MapMany(productEntities *[]entities.ProductEntity) *[]models.ProductModel {
	productModels := make([]models.ProductModel, 0)

	for _, productEntity := range *productEntities {
		productModel := m.MapOne(&productEntity)
		productModels = append(productModels, *productModel)
	}

	return &productModels
}

func (m ProductMapper) MapOne(productEntity *entities.ProductEntity) *models.ProductModel {
	return &models.ProductModel{
		Id:        productEntity.Id,
		Name:      productEntity.Name,
		Price:     productEntity.Price,
		Status:    productEntity.Status,
		CreatedAt: productEntity.CreatedAt.Unix(),
		UpdatedAt: productEntity.UpdatedAt.Unix(),
		DeletedAt: productEntity.DeletedAt.Unix(),
	}
}
