package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"cli-sample/internal/core/models"
)

type PriceService struct {
}

func NewPriceService() *PriceService {
	return &PriceService{}
}

func (s *PriceService) GetPrice(symbol string) (*models.PriceModel, error) {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=" + symbol)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var price models.PriceModel

	err = json.Unmarshal(body, &price)
	if err != nil {
		return nil, err
	}

	return &price, nil
}
