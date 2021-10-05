package main

import (
	"context"
	"fmt"
	"math/rand"

	"client-sample/pkg/clients"
	"client-sample/pkg/configs"
	"client-sample/pkg/constants"
	"client-sample/pkg/requests"
)

var productClient *clients.ProductClient

func main() {
	productClient = clients.NewProductClient(&configs.ProductClientConfig{
		BaseUrl:  "http://localhost:2805",
		Username: "admin",
		Password: "123456",
	})

	id := CreateProduct()
	GetProducts()
	GetProductById(id)
	UpdateProduct(id)
	DeleteProduct(id)
}

func GetProducts() {
	getProductsResponse, err := productClient.GetProducts(context.Background(), &requests.GetProductsRequest{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(getProductsResponse.Products)
}

func GetProductById(id string) {
	getProductByIdResponse, err := productClient.GetProductById(context.Background(), &requests.GetProductByIdRequest{
		Id: id,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(getProductByIdResponse.Product)
}

func CreateProduct() string {
	createProductResponse, err := productClient.CreateProduct(context.Background(), &requests.CreateProductRequest{
		Name:   fmt.Sprintf("product %v", rand.Intn(10)),
		Price:  50,
		Status: constants.ProductStatusEnabled,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(createProductResponse.Id)

	return createProductResponse.Id
}

func UpdateProduct(id string) {
	updateProductResponse, err := productClient.UpdateProduct(context.Background(), &requests.UpdateProductRequest{
		Id:     id,
		Name:   fmt.Sprintf("product %v", rand.Intn(10)),
		Price:  50,
		Status: constants.ProductStatusDisabled,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(updateProductResponse)
}

func DeleteProduct(id string) {
	deleteProductResponse, err := productClient.DeleteProduct(context.Background(), &requests.DeleteProductRequest{
		Id: id,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteProductResponse)
}
