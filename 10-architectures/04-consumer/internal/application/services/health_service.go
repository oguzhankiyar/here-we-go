package services

import (
	healthModels "consumer-sample/internal/core/models"
	configModels "consumer-sample/internal/infrastructure/config/models"
	"consumer-sample/internal/infrastructure/logger/interfaces"
	"time"
)

type HealthService struct {
	config  configModels.AppConfig
	logger  interfaces.Logger
	history map[string]healthModels.HistoryModel
}

func NewHealthService(config configModels.AppConfig, logger interfaces.Logger) *HealthService {
	return &HealthService{
		config:  config,
		logger:  logger,
		history: make(map[string]healthModels.HistoryModel),
	}
}

func (s *HealthService) Set(model *healthModels.HealthModel) {
	s.logger.Debug(model.App + " " + model.State)

	if h, ok := s.history[model.App]; ok {
		if h.State == "unhealthy" && model.State == "health" {
			s.logger.Info(model.App + ": unhealthy -> healthy")
		} else if h.State == "healthy" && model.State == "unhealthy" {
			s.logger.Info(model.App + ": healthy -> unhealthy")
		}
	} else {
		s.logger.Info(model.App + ": unknown -> " + model.State)
	}

	s.history[model.App] = healthModels.HistoryModel{
		Time:  model.Time,
		State: model.State,
	}
}

func (s *HealthService) Check() {
	for a, h := range s.history {
		now := time.Now().Unix()
		diff := now - h.Time
		if diff > s.config.HealthTreshold {
			s.history[a] = healthModels.HistoryModel{
				Time:  now,
				State: "unhealthy",
			}

			s.logger.Info(a + ": healthy -> unhealthy")
		}
	}
}
