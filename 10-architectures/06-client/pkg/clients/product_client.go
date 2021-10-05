package clients

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"

	"client-sample/pkg/configs"
	"client-sample/pkg/models"
	"client-sample/pkg/requests"
	"client-sample/pkg/responses"
)

type ProductClient struct {
	client *resty.Client
	config *configs.ProductClientConfig
}

func NewProductClient(config *configs.ProductClientConfig) *ProductClient {
	client := resty.New()
	client.
		SetRetryCount(2).
		SetRetryWaitTime(500 * time.Millisecond).
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryHook(func(response *resty.Response, err error) {
			fmt.Println(time.Now(), "retrying")
		}).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() == 500
			},
		)

	return &ProductClient{
		client: client,
		config: config,
	}
}

func (c *ProductClient) GetProducts(ctx context.Context, request *requests.GetProductsRequest) (*responses.GetProductsResponse, error) {
	url := fmt.Sprintf("%s/products", c.config.BaseUrl)

	token, err := c.getAuthToken(ctx)
	if err != nil {
		return nil, err
	}

	var response models.BaseReponseModel

	_, err = c.client.R().
		SetQueryParams(map[string]string{
			"offset": strconv.Itoa(request.Offset),
			"limit":  strconv.Itoa(request.Limit),
			"sort":   request.Sort,
			"order":  strconv.Itoa(int(request.Order)),
		}).
		SetHeader("Accept", "application/json").
		SetAuthToken(token.AccessToken).
		SetResult(&response).
		Get(url)

	if err != nil {
		return nil, err
	}

	if response.Data == nil {
		return nil, errors.New(response.Code + " - " + response.Message)
	}

	var products []models.ProductModel

	err = mapstructure.Decode(response.Data, &products)

	if err != nil {
		return nil, err
	}

	return &responses.GetProductsResponse{
		Products: &products,
	}, nil
}

func (c *ProductClient) GetProductById(ctx context.Context, request *requests.GetProductByIdRequest) (*responses.GetProductByIdResponse, error) {
	url := fmt.Sprintf("%s/products/%s", c.config.BaseUrl, request.Id)

	token, err := c.getAuthToken(ctx)
	if err != nil {
		return nil, err
	}

	var response models.BaseReponseModel

	_, err = c.client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(token.AccessToken).
		SetResult(&response).
		Get(url)

	if err != nil {
		return nil, err
	}

	if response.Data == nil {
		return nil, errors.New(response.Code + " - " + response.Message)
	}

	var product models.ProductModel

	err = mapstructure.Decode(response.Data, &product)

	if err != nil {
		return nil, err
	}

	return &responses.GetProductByIdResponse{
		Product: &product,
	}, nil
}

func (c *ProductClient) CreateProduct(ctx context.Context, request *requests.CreateProductRequest) (*responses.CreateProductResponse, error) {
	url := fmt.Sprintf("%s/products", c.config.BaseUrl)

	token, err := c.getAuthToken(ctx)
	if err != nil {
		return nil, err
	}

	var response models.BaseReponseModel

	_, err = c.client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(token.AccessToken).
		SetBody(request).
		SetResult(&response).
		Post(url)

	if err != nil {
		return nil, err
	}

	if response.Data == nil {
		return nil, errors.New(response.Code + " - " + response.Message)
	}

	var id string

	err = mapstructure.Decode(response.Data, &id)

	if err != nil {
		return nil, err
	}

	return &responses.CreateProductResponse{
		Id: id,
	}, nil
}

func (c *ProductClient) UpdateProduct(ctx context.Context, request *requests.UpdateProductRequest) (*responses.UpdateProductResponse, error) {
	url := fmt.Sprintf("%s/products/%s", c.config.BaseUrl, request.Id)

	token, err := c.getAuthToken(ctx)
	if err != nil {
		return nil, err
	}

	response, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(token.AccessToken).
		SetBody(request).
		Put(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 204 {
		return nil, errors.New("failed")
	}

	return &responses.UpdateProductResponse{}, nil
}

func (c *ProductClient) DeleteProduct(ctx context.Context, request *requests.DeleteProductRequest) (*responses.DeleteProductResponse, error) {
	url := fmt.Sprintf("%s/products/%s", c.config.BaseUrl, request.Id)

	token, err := c.getAuthToken(ctx)
	if err != nil {
		return nil, err
	}

	response, err := c.client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken(token.AccessToken).
		Delete(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 204 {
		return nil, errors.New("failed")
	}

	return &responses.DeleteProductResponse{}, nil
}

func (c *ProductClient) getAuthToken(ctx context.Context) (*models.TokenModel, error) {
	url := fmt.Sprintf("%s/auth/token", c.config.BaseUrl)

	var response models.BaseReponseModel

	_, err := c.client.R().
		SetFormData(map[string]string{
			"username": c.config.Username,
			"password": c.config.Password,
		}).
		SetHeader("Accept", "application/json").
		SetResult(&response).
		Post(url)

	if err != nil {
		return nil, err
	}

	if response.Data == nil {
		return nil, errors.New(response.Code + " - " + response.Message)
	}

	var token models.TokenModel

	err = mapstructure.Decode(response.Data, &token)

	if err != nil {
		return nil, err
	}

	return &token, nil
}
