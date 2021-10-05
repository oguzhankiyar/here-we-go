package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"web-sample/internal/common/constants"
	"web-sample/internal/core/product/entities"
	"web-sample/internal/core/product/repositories"
	"web-sample/internal/infrastructure/persistence"
)

type productRepository struct {
	postgresPersistence *persistence.PostgresPersistence
}

func NewProductRepository(postgresPersistence *persistence.PostgresPersistence) repositories.ProductRepository {
	return &productRepository{
		postgresPersistence: postgresPersistence,
	}
}

func (r *productRepository) Find(ctx context.Context, offset, limit int, sort string, order constants.Order) (*[]entities.ProductEntity, error) {
	var products []entities.ProductEntity

	if len(sort) == 0 {
		sort = "created_at"
	}

	if order == constants.OrderDescending {
		sort += " desc"
	} else {
		sort += " asc"
	}

	result := r.postgresPersistence.Database.
		Offset(offset).
		Limit(limit).
		Order(sort).
		Find(&products, "is_deleted = ?", false)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (r *productRepository) FindById(ctx context.Context, id string) (*entities.ProductEntity, error) {
	var product entities.ProductEntity

	result := r.postgresPersistence.Database.
		First(&product, "is_deleted = ? AND id = ?", false, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &product, nil
}

func (r *productRepository) FindByName(ctx context.Context, name string) (*entities.ProductEntity, error) {
	var product entities.ProductEntity

	result := r.postgresPersistence.Database.
		First(&product, "is_deleted = ? AND name = ?", false, name)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &product, nil
}

func (r *productRepository) Create(ctx context.Context, product *entities.ProductEntity) (string, error) {
	product.Id = uuid.New().String()
	product.CreatedAt = time.Now()

	result := r.postgresPersistence.Database.
		Create(&product)
	if result.Error != nil {
		return "", result.Error
	}

	return product.Id, nil
}

func (r *productRepository) Update(ctx context.Context, product *entities.ProductEntity) error {
	product.UpdatedAt = time.Now()

	result := r.postgresPersistence.Database.
		Save(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *productRepository) Delete(ctx context.Context, id string) error {
	var product entities.ProductEntity

	result := r.postgresPersistence.Database.
		First(&product, "id = ?", id)
	if result.Error != nil {
		return nil
	}

	product.IsDeleted = true
	product.DeletedAt = time.Now()

	result = r.postgresPersistence.Database.
		Save(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
