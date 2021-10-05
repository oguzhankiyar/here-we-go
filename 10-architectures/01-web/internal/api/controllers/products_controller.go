package controllers

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"web-sample/internal/application/handlers/product"
	"web-sample/internal/common/models"
	"web-sample/internal/core/product/commands"
	"web-sample/internal/core/product/queries"
)

type ProductsController struct {
	productHandler *product.ProductHandler
	validator      *validator.Validate
}

func NewProductsController(productHandler *product.ProductHandler, validator *validator.Validate) *ProductsController {
	return &ProductsController{
		productHandler: productHandler,
		validator:      validator,
	}
}

// HandleGetAll godoc
// @Summary Get all products.
// @Description get all products.
// @Tags Products
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.BaseReponseModel
// @Failure 500 {object} models.BaseReponseModel
// @Router /products [get]
func (c *ProductsController) HandleGetAll(ctx echo.Context) error {
	qry := new(queries.FindProductsQuery)
	if err := ctx.Bind(&qry); err != nil {
		return err
	}

	if err := c.validator.Struct(qry); err != nil {
		return err
	}

	products, err := c.productHandler.FindProducts.Handle(context.Background(), qry)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.BaseReponseModel{
		Code:    "200",
		Message: "SUCCEEDED",
		Errors:  []string{},
		Data:    products,
	})
}

// HandleGetById godoc
// @Summary Get product by id.
// @Description get product by id.
// @Tags Products
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.BaseReponseModel
// @Failure 404 {object} models.BaseReponseModel
// @Failure 500 {object} models.BaseReponseModel
// @Router /products/{id} [get]
func (c *ProductsController) HandleGetById(ctx echo.Context) error {
	qry := new(queries.FindProductByIdQuery)
	qry.Id = ctx.Param("id")

	if err := c.validator.Struct(qry); err != nil {
		return err
	}

	product, err := c.productHandler.FindProductById.Handle(context.Background(), qry)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.BaseReponseModel{
		Code:    "200",
		Message: "SUCCEEDED",
		Errors:  []string{},
		Data:    product,
	})
}

// HandleCreate godoc
// @Summary Create product.
// @Description create product.
// @Tags Products
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Param productData body commands.CreateProductCommand true "product data"
// @Success 201 {object} models.BaseReponseModel
// @Failure 409 {object} models.BaseReponseModel
// @Failure 500 {object} models.BaseReponseModel
// @Router /products [post]
func (c *ProductsController) HandleCreate(ctx echo.Context) error {
	cmd := new(commands.CreateProductCommand)
	if err := ctx.Bind(&cmd); err != nil {
		return err
	}

	if err := c.validator.Struct(cmd); err != nil {
		return err
	}

	id, err := c.productHandler.CreateProduct.Handle(context.Background(), cmd)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, models.BaseReponseModel{
		Code:    "201",
		Message: "SUCCEEDED",
		Errors:  []string{},
		Data:    id,
	})
}

// HandleUpdate godoc
// @Summary Update product by id.
// @Description update product by id.
// @Tags Products
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "product id"
// @Success 204
// @Failure 404 {object} models.BaseReponseModel
// @Failure 500 {object} models.BaseReponseModel
// @Router /products/{id} [put]
func (c *ProductsController) HandleUpdate(ctx echo.Context) error {
	cmd := new(commands.UpdateProductCommand)
	if err := ctx.Bind(&cmd); err != nil {
		return err
	}
	cmd.Id = ctx.Param("id")

	if err := c.validator.Struct(cmd); err != nil {
		return err
	}

	err := c.productHandler.UpdateProduct.Handle(context.Background(), cmd)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// HandleDelete godoc
// @Summary Delete product by id.
// @Description delete product by id.
// @Tags Products
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "product id"
// @Success 204
// @Failure 404 {object} models.BaseReponseModel
// @Failure 500 {object} models.BaseReponseModel
// @Router /products/{id} [delete]
func (c *ProductsController) HandleDelete(ctx echo.Context) error {
	cmd := new(commands.DeleteProductCommand)
	cmd.Id = ctx.Param("id")

	if err := c.validator.Struct(cmd); err != nil {
		return err
	}

	err := c.productHandler.DeleteProduct.Handle(context.Background(), cmd)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
