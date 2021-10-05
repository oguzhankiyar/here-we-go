package errors

import "web-sample/internal/infrastructure/errors"

const (
	ProductNotFoundErrorCode     = "310"
	ProductAlreadyExistErrorCode = "311"
)

const (
	ProductNotFoundErrorMessage     = "product could not found"
	ProductAlreadyExistErrorMessage = "product is already exist"
)

var (
	ProductNotFoundError     = errors.NewNotFoundError(ProductNotFoundErrorCode, ProductNotFoundErrorMessage)
	ProductAlreadyExistError = errors.NewAlreadyExistError(ProductAlreadyExistErrorCode, ProductAlreadyExistErrorMessage)
)
