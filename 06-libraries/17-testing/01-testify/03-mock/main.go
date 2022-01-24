package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
)

type ServiceInterface interface {
	GetCounts(length int) (map[string]int, error)
}

type Service struct {

}

func (s *Service) GetCounts(length int) (map[string]int, error) {
	if length <= 0 {
		return nil, errors.New("the length should be greater than zero")
	}

	result := make(map[string]int)
	for i := 0; i < length; i++ {
		id := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", rand.Int())))[0:10]
		count := rand.Int() % 100 + 100
		result[id] = count
	}
	return result, nil
}

func Run(service ServiceInterface, id string) (int, error) {
	if len(id) == 0 {
		return 0, errors.New("the id could not be empty")
	}

	result, err := service.GetCounts(10)

	if err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, errors.New("the id is not found")
	}

	if count, ok := result[id]; ok {
		return count, nil
	}

	return 0, errors.New("the id is not found")
}
